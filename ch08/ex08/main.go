// Copyright 2017 Ken Miura
// reverb2を修正して作成
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	ch := make(chan string)
	go func() {
		for input.Scan() {
			ch <- input.Text()
		}
		close(ch)
	}()
receive:
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Fprintln(c, "server stopped receiving due to no requet for 10 seconds")
			break receive
		case text, ok := <-ch:
			if !ok {
				break receive
			}
			go echo(c, text, 1*time.Second)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
	for range ch {
		// do nothing // ch<-input.Text()でブロックされてゴルーチンのリークにならないようにループする。
	}
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
