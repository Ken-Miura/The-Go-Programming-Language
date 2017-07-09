// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func handleConn(c net.Conn) {
	defer c.Close()
	c.Write([]byte("220 Service ready for new user.\n"))
	clientIP, portString, err := net.SplitHostPort(c.RemoteAddr().String())
	if err != nil {
		log.Print("cannot get client IP address")
		clientIP = ""
	}
	tmp, err := strconv.ParseInt(portString, 10, 0)
	clientPortForDataTransfer := int(tmp)
	if err != nil {
		log.Print("cannot get default port for data transfer to client")
		clientPortForDataTransfer = -1
	}
	input := bufio.NewScanner(c)
	for input.Scan() {
		if err := input.Err(); err != nil {
			// 適切なステータスコード打ち込んで返す TODO
			continue
		}
		line := input.Text()
		log.Print(line)
		commandAndArgs := strings.Fields(line)
		command := strings.ToUpper(strings.ToLower(commandAndArgs[0]))
		args := commandAndArgs[1:]
		switch command {
		case "USER":
			c.Write([]byte(fmt.Sprintf("230 User logged in, proceed. response for command (%s)\n", line)))
			// TODO
			// パスワード要求 → 331
			// ログイン成功 → 230
			// ログイン失敗 → 530
		case "PASS":
			c.Write([]byte(fmt.Sprintf("202 Command not implemented, superfluous at this site. response for command (%s)\n", line)))
		case "ACCT":
			c.Write([]byte(fmt.Sprintf("202 Command not implemented, superfluous at this site. response for command (%s)\n", line)))
		case "QUIT":
			/* TODO
			 * データ転送がない場合、コントロール接続を閉じる。
			 * データ転送がある場合、転送結果を送信後にコントロール接続を閉じる。
			 * コントロール接続が予期せず終了した場合、サーバーは中断(ABOR)とログアウト(QUIT)との効果を持つ動作
			 */
			c.Write([]byte(fmt.Sprintf("221 Service closing control connection. response for command (%s)\n", line)))
		case "PORT": // IPv6はEPRTコマンドで渡されてくるので、このコマンドの処理はIPv4を想定したものでOK
			/* TODO
			 * put系操作の場合
			 * 次に125 Data connection already open; transfer starting.
			 * 最後に226
			 * get系操作の場合
			 * 流れ的にはputと同じ
			 */
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			clientIP, clientPortForDataTransfer = port(c, args[0], line)
		case "TYPE":
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		case "MODE":
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		case "STRU":
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		case "ALLO":
			c.Write([]byte(fmt.Sprintf("202 Command not implemented, superfluous at this site. response for command (%s)\n", line)))
		case "RETR":
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		case "STOR":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			stor(c, args[0], clientIP, clientPortForDataTransfer, line)
		case "STOU":
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		case "SITE":
			c.Write([]byte(fmt.Sprintf("202 Command not implemented, superfluous at this site. response for command (%s)\n", line)))
		case "NOOP":
			c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
		default:
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		}
	}
}

func port(c net.Conn, arg, line string) (string, int) {
	IPv4AndPort := strings.Split(arg, ",")
	clientIP := IPv4AndPort[0] + "." + IPv4AndPort[1] + "." + IPv4AndPort[2] + "." + IPv4AndPort[3]
	firstSegmentOfPortNum, err := strconv.ParseInt(IPv4AndPort[4], 10, 0)
	if err != nil {
		c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	secondSegmentOfPortNum, err := strconv.ParseInt(IPv4AndPort[5], 10, 0)
	if err != nil {
		c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	clientPortForDataTransfer := int(firstSegmentOfPortNum)*256 + int(secondSegmentOfPortNum)
	if clientPortForDataTransfer < 0 || clientPortForDataTransfer > 65535 {
		c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
	return clientIP, clientPortForDataTransfer
}

func stor(c net.Conn, fileName string, clientIP string, clientPort int, line string) {
	f, err := os.Create(fileName)
	if err != nil {
		c.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	defer f.Close()
	d := net.Dialer{LocalAddr: &net.TCPAddr{IP: net.ParseIP(*ip), Port: *portForControlConnection - 1}}
	connForDataTransfer, err := d.Dial("tcp", fmt.Sprintf("%s:%d", clientIP, clientPort))
	if err != nil {
		fmt.Println(err)
		c.Write([]byte(fmt.Sprintf("425 Can't open data connection. response for command (%s)\n", line)))
		return
	}
	defer connForDataTransfer.Close()
	c.Write([]byte(fmt.Sprintf("125 Data connection already open; transfer starting. response for command (%s)\n", line)))
	_, err = io.Copy(f, connForDataTransfer)
	if err != nil {
		c.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	c.Write([]byte(fmt.Sprintf("226 Closing data connection. Requested file action successful. response for command (%s)\n", line)))
}

var ip = flag.String("ip", "localhost", "IP address for binding")
var portForControlConnection = flag.Int("port", 21, "port number for control connection ")

func init() {
	flag.Parse()

}

func main() {
	if *portForControlConnection < 0 {
		fmt.Println("Port number must be 0 or more.")
		return
	}
	if *ip == "" {
		fmt.Println("IP address must not be empty.")
		return
	}
	if strings.Contains(*ip, ":") { // IPv6のとき[]で囲む必要があるため、IPv6かどうかの判定
		// TODO IPv6をサポートするためにはEPRTコマンドに対応しなければならない（RFC2428)
		//*ip = "[" + *ip + "]" net.JoinHostPort実装の際の参考
		fmt.Println("IPv6 is not supported. Use IPv4 instead.")
		return
	}

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *portForControlConnection))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
