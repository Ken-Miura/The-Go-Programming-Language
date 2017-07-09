// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConn(c net.Conn) {
	c.Write([]byte("200 Service ready for new user.\n"))
	input := bufio.NewScanner(c)
	for input.Scan() {
		if err := input.Err(); err != nil {
			// 適切なステータスコード打ち込んで返す TODO
			// c.Write([]byte(""))
			continue
		}
		line := input.Text()
		fmt.Println(line) // debug用にプリント
		commandAndArgs := strings.Fields(line)
		command := strings.ToUpper(strings.ToLower(commandAndArgs[0]))
		//args := commandAndArgs[1:]
		switch command {
		case "USER":
			c.Write([]byte("502 Command not implemented.\n"))
		case "QUIT":
			c.Write([]byte("502 Command not implemented.\n"))
		case "PORT":
			c.Write([]byte("502 Command not implemented.\n"))
		case "TYPE":
			c.Write([]byte("502 Command not implemented.\n"))
		case "MODE":
			c.Write([]byte("502 Command not implemented.\n"))
		case "STRU":
			c.Write([]byte("502 Command not implemented.\n"))
		case "RETR":
			c.Write([]byte("502 Command not implemented.\n"))
		case "STOR":
			c.Write([]byte("502 Command not implemented.\n"))
		case "NOOP":
			c.Write([]byte("502 Command not implemented.\n"))
		default:
			c.Write([]byte("502 Command not implemented.\n"))
		}
	}
	c.Close()
}

var port = flag.Int("port", 21, "port number")
var ip = flag.String("ip", "localhost", "IP address for binding")

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
		*ip = "[" + *ip + "]"
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
