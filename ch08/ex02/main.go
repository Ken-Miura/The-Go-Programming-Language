// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func handleConn(c net.Conn) {
	c.Write([]byte("220 Service ready for new user.\n"))
	//portForDataTransfer := *port - 1
	clientIP := ""
	fmt.Print(clientIP)
	clientPortForDataTransfer := -1
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
			 * 成功 → 200
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
			IPv4AndPort := strings.Split(args[0], ",")
			clientIP = IPv4AndPort[0] + "." + IPv4AndPort[1] + "." + IPv4AndPort[2] + "." + IPv4AndPort[3]
			firstByteOfPortNum, err := strconv.ParseInt(IPv4AndPort[4], 10, 0)
			if err != nil {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			secondByteOfPortNum, err := strconv.ParseInt(IPv4AndPort[5], 10, 0)
			if err != nil {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			clientPortForDataTransfer = int(firstByteOfPortNum)*4 + int(secondByteOfPortNum)
			if clientPortForDataTransfer < 0 || clientPortForDataTransfer > 65535 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
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
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
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
	c.Close()
}

var ip = flag.String("ip", "localhost", "IP address for binding")
var port = flag.Int("port", 21, "port number for control connection ")

func init() {
	flag.Parse()

}

func main() {
	if *port < 0 {
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

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *ip, *port))
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
