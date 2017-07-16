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
	event := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Fprintln(c, "server stopped receiving due to no requet for 10 seconds")
				c.Close()
				for range event {
					// do nothing
				}
				return
			case _, ok := <-event:
				if !ok {
					return
				}
			}
		}
	}()
	for input.Scan() {
		event <- struct{}{}
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
	close(event)
	c.Close()
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
