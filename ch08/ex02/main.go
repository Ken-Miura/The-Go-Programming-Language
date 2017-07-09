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
	"sync"
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
	var wg sync.WaitGroup
	for input.Scan() {
		if err := input.Err(); err != nil {
			c.Write([]byte("500 Syntax error, command unrecognized.\n"))
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
			wg.Wait()
			c.Write([]byte(fmt.Sprintf("221 Service closing control connection. response for command (%s)\n", line)))
		case "PORT": // IPv6はEPRTコマンドで渡されてくるので、このコマンドの処理はIPv4を想定したものでOK
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			clientIP, clientPortForDataTransfer = port(c, args[0], line)
		case "TYPE":
			c.Write([]byte(fmt.Sprintf("200 Command not implemented. response for command (%s)\n", line)))
		case "MODE":
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		case "STRU":
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		case "ALLO":
			c.Write([]byte(fmt.Sprintf("202 Command not implemented, superfluous at this site. response for command (%s)\n", line)))
		case "RETR":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				retr(c, args[0], clientIP, clientPortForDataTransfer, line)
			}()
		case "STOR":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				stor(c, args[0], clientIP, clientPortForDataTransfer, line)
			}()
		case "STOU":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				stou(c, args[0], clientIP, clientPortForDataTransfer, line)
			}()
		case "SITE":
			c.Write([]byte(fmt.Sprintf("202 Command not implemented, superfluous at this site. response for command (%s)\n", line)))
		case "NOOP":
			c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
		default:
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		}
	}
}

func port(out io.Writer, arg, line string) (string, int) {
	IPv4AndPort := strings.Split(arg, ",")
	clientIP := IPv4AndPort[0] + "." + IPv4AndPort[1] + "." + IPv4AndPort[2] + "." + IPv4AndPort[3]
	firstSegmentOfPortNum, err := strconv.ParseInt(IPv4AndPort[4], 10, 0)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	secondSegmentOfPortNum, err := strconv.ParseInt(IPv4AndPort[5], 10, 0)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	clientPortForDataTransfer := int(firstSegmentOfPortNum)*256 + int(secondSegmentOfPortNum)
	if clientPortForDataTransfer < 0 || clientPortForDataTransfer > 65535 {
		out.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	out.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
	return clientIP, clientPortForDataTransfer
}

func stor(out io.Writer, fileName string, clientIP string, clientPort int, line string) {
	f, err := os.Create(fileName)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	defer f.Close()
	d := net.Dialer{LocalAddr: &net.TCPAddr{IP: net.ParseIP(*ip), Port: *portForControlConnection - 1}}
	connForDataTransfer, err := d.Dial("tcp", fmt.Sprintf("%s:%d", clientIP, clientPort))
	if err != nil {
		fmt.Println(err)
		out.Write([]byte(fmt.Sprintf("425 Can't open data connection. response for command (%s)\n", line)))
		return
	}
	defer connForDataTransfer.Close()
	transferData(out, f, connForDataTransfer, "stored "+fileName, line)
}

func retr(out io.Writer, fileName string, clientIP string, clientPort int, line string) {
	f, err := os.Open(fileName)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	defer f.Close()
	d := net.Dialer{LocalAddr: &net.TCPAddr{IP: net.ParseIP(*ip), Port: *portForControlConnection - 1}}
	connForDataTransfer, err := d.Dial("tcp", fmt.Sprintf("%s:%d", clientIP, clientPort))
	if err != nil {
		fmt.Println(err)
		out.Write([]byte(fmt.Sprintf("425 Can't open data connection. response for command (%s)\n", line)))
		return
	}
	defer connForDataTransfer.Close()
	transferData(out, connForDataTransfer, f, "retrieved "+fileName, line)
}

func stou(out io.Writer, fileName string, clientIP string, clientPort int, line string) {
	for i := 0; true; i++ {
		if _, err := os.Stat(fileName); err != nil {
			break
		}
		fileName = fileName + fmt.Sprintf(".%d", i) // すでに指定されたファイル名が存在するため、一意なファイル名を用意
	}
	f, err := os.Create(fileName)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	defer f.Close()
	d := net.Dialer{LocalAddr: &net.TCPAddr{IP: net.ParseIP(*ip), Port: *portForControlConnection - 1}}
	connForDataTransfer, err := d.Dial("tcp", fmt.Sprintf("%s:%d", clientIP, clientPort))
	if err != nil {
		fmt.Println(err)
		out.Write([]byte(fmt.Sprintf("425 Can't open data connection. response for command (%s)\n", line)))
		return
	}
	defer connForDataTransfer.Close()
	transferData(out, f, connForDataTransfer, "stored "+fileName, line)
}

func transferData(out io.Writer, dst io.Writer, src io.Reader, message, line string) {
	out.Write([]byte(fmt.Sprintf("125 Data connection already open; transfer starting. response for command (%s)\n", line)))
	_, err := io.Copy(dst, src)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	out.Write([]byte(fmt.Sprintf("250 Requested file action okay, completed (%s). response for command (%s)\n", message, line)))
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
