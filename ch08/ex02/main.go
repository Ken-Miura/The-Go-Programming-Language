// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var ip = flag.String("ip", "", "IP address for binding")
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
		fmt.Println("usage: " + os.Args[0] + " -ip 'IPv4 address'")
		fmt.Println("ex: " + os.Args[0] + " -ip 192.168.0.5")
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

type dataType int

const (
	ASCII dataType = iota
	IMAGE
)

type dataStructure int

const (
	FILE dataStructure = iota
	RECORD
)

func handleConn(c net.Conn) {
	defer c.Close()
	c.Write([]byte("220 Service ready for new user.\n"))
	clientIP, portString, err := net.SplitHostPort(c.RemoteAddr().String())
	dataType := ASCII
	dataStructure := FILE
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
			if !(len(args) == 1 || len(args) == 2) {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			if args[0] == "A" {
				dataType = ASCII
				if len(args) == 2 && args[0] != "NON PRINT" {
					c.Write([]byte(fmt.Sprintf("200 We support only NON PRINT. response for command (%s)\n", line)))
				} else {
					c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
				}
			} else if args[0] == "I" {
				dataType = IMAGE
				c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
			} else {
				c.Write([]byte(fmt.Sprintf("200 We support only either ASCII TYPE or IMAGE TYPE. response for command (%s)\n", line)))
			}
		case "MODE":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			if args[0] != "S" {
				c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
			} else {
				c.Write([]byte(fmt.Sprintf("200 We support only stream mode. response for command (%s)\n", line)))
			}
		case "STRU":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			if args[0] == "F" {
				dataStructure = FILE
				c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
			} else if args[0] == "R" {
				dataStructure = RECORD
				c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
			} else {
				c.Write([]byte(fmt.Sprintf("200 We support only either file-structure or record-structure. response for command (%s)\n", line)))
			}
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
				retr(c, args[0], clientIP, clientPortForDataTransfer, dataType, dataStructure, line)
			}()
		case "STOR":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				stor(c, args[0], clientIP, clientPortForDataTransfer, dataType, dataStructure, line)
			}()
		case "STOU":
			if len(args) != 1 {
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				stou(c, args[0], clientIP, clientPortForDataTransfer, dataType, dataStructure, line)
			}()
		case "SITE":
			c.Write([]byte(fmt.Sprintf("202 Command not implemented, superfluous at this site. response for command (%s)\n", line)))
		case "NOOP":
			c.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
		case "LIST":
			if len(args) != 0 { // ファイル名やディレクトリ名が引数で指定されてくる場合がある。しかし現在は未サポートで、ワーキングディレクトリの情報一覧を取ってくる処理のみサポート。
				c.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
				break
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				list(c, ".", clientIP, clientPortForDataTransfer, dataType, dataStructure, line)
			}()

		default:
			c.Write([]byte(fmt.Sprintf("502 Command not implemented. response for command (%s)\n", line)))
		}
	}
}

func port(out io.Writer, arg, line string) (string, int) {
	IPv4AndPort := strings.Split(arg, ",")
	clientIP := IPv4AndPort[0] + "." + IPv4AndPort[1] + "." + IPv4AndPort[2] + "." + IPv4AndPort[3]
	first, err := strconv.ParseInt(IPv4AndPort[4], 10, 0)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	second, err := strconv.ParseInt(IPv4AndPort[5], 10, 0)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	clientPortForDataTransfer := int(first)*256 + int(second)
	if clientPortForDataTransfer < 0 || clientPortForDataTransfer > 65535 {
		out.Write([]byte(fmt.Sprintf("501 Syntax error in parameters or arguments. response for command (%s)\n", line)))
		return "", -1
	}
	out.Write([]byte(fmt.Sprintf("200 Command okay. response for command (%s)\n", line)))
	return clientIP, clientPortForDataTransfer
}

func stor(out io.Writer, fileName string, clientIP string, clientPort int, dataType dataType, dataStructure dataStructure, line string) {
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
	transferData(out, f, connForDataTransfer, dataType, dataStructure, "stored "+fileName, line)
}

func retr(out io.Writer, fileName string, clientIP string, clientPort int, dataType dataType, dataStructure dataStructure, line string) {
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
	transferData(out, connForDataTransfer, f, dataType, dataStructure, "retrieved "+fileName, line)
}

func stou(out io.Writer, fileName string, clientIP string, clientPort int, dataType dataType, dataStructure dataStructure, line string) {
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
	transferData(out, f, connForDataTransfer, dataType, dataStructure, "stored "+fileName, line)
}

func list(out io.Writer, item string, clientIP string, clientPort int, dataType dataType, dataStructure dataStructure, line string) {
	itemInfo, err := os.Stat(item)
	if err != nil {
		out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	var buf bytes.Buffer
	if itemInfo.IsDir() {
		filesInfo, err := ioutil.ReadDir(item)
		if err != nil {
			out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
			return
		}
		for _, fileInfo := range filesInfo {
			buf.WriteString(fileInfo.ModTime().String())
			buf.WriteString("\t")
			if fileInfo.IsDir() {
				buf.WriteString("<DIR>")
				buf.WriteString("\t")
				buf.WriteString("\t")
			} else {
				buf.WriteString("\t")
				buf.WriteString(strconv.FormatInt(fileInfo.Size(), 10))
				buf.WriteString("\t")
			}
			buf.WriteString(fileInfo.Name())
			buf.WriteString("\n")
		}
	} else {
		buf.WriteString(itemInfo.ModTime().String())
		buf.WriteString("\t")
		buf.WriteString("\t")
		buf.WriteString(strconv.FormatInt(itemInfo.Size(), 10))
		buf.WriteString("\t")
		buf.WriteString(itemInfo.Name())
		buf.WriteString("\n")
	}

	d := net.Dialer{LocalAddr: &net.TCPAddr{IP: net.ParseIP(*ip), Port: *portForControlConnection - 1}}
	connForDataTransfer, err := d.Dial("tcp", fmt.Sprintf("%s:%d", clientIP, clientPort))
	if err != nil {
		fmt.Println(err)
		out.Write([]byte(fmt.Sprintf("425 Can't open data connection. response for command (%s)\n", line)))
		return
	}
	defer connForDataTransfer.Close()
	transferData(out, connForDataTransfer, buf, dataType, dataStructure, "list "+itemInfo.Name(), line)
}

func transferData(out io.Writer, dst io.Writer, src io.Reader, dataType dataType, dataStructure dataStructure, message, line string) {
	out.Write([]byte(fmt.Sprintf("125 Data connection already open; transfer starting. response for command (%s)\n", line)))
	var err error
	if dataType == ASCII && dataStructure == RECORD {
		var b [4096]byte
		for {
			n, err := src.Read(b[:])
			if n > 0 {
				var writeErr error
				str := string(b[:n])
				if strings.Contains(str, "\r") || strings.Contains(str, "\n") {
					// \r → \r\n
					var bufForCR bytes.Buffer
					for i := 0; i < n; i++ {
						bufForCR.WriteByte(str[i])
						if str[i] == '\r' && i != n-1 && str[i+1] != byte('\n') {
							bufForCR.WriteByte('\n')
						}
					}
					// \n → \r\n
					var bufForLF bytes.Buffer
					tmp := string(bufForCR.Bytes())
					for i := 0; i < len(tmp); i++ {
						if tmp[i] == '\n' && i != 0 && tmp[i-1] != '\r' {
							bufForLF.WriteByte('\r')
						}
						bufForLF.WriteByte(tmp[i])
					}
					_, writeErr = dst.Write(bufForLF.Bytes())
				} else {
					_, writeErr = dst.Write(b[:n])
				}
				if err == nil {
					err = writeErr
				}
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
				return
			}
		}
	} else {
		_, err = io.Copy(dst, src)
	}
	if err != nil {
		out.Write([]byte(fmt.Sprintf("451 Requested action aborted. response for command (%s)\n", line)))
		return
	}
	out.Write([]byte(fmt.Sprintf("250 Requested file action okay, completed (%s). response for command (%s)\n", message, line)))
}
