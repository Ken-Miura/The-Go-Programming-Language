// Copyright 2017 Ken Miura
// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			echo(c, input.Text(), 1*time.Second)
		}()
	}
	// NOTE: ignoring potential errors from input.Err()
	tcpConn := c.(*net.TCPConn) // 練習問題３の回答とセットで実施するため、*net.TCPConn以外ありえないので失敗したらpanicで終了するようにする。
	tcpConn.CloseRead()
	wg.Wait()
	tcpConn = c.(*net.TCPConn) // 練習問題３の回答とセットで実施するため、*net.TCPConn以外ありえないので失敗したらpanicで終了するようにする。
	tcpConn.CloseWrite()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
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
