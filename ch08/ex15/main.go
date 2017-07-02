// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client struct {
	name string
	ch   chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
					// クライアント側で書き込み準備ができていないとclientsの他のクライアントのすべてが待たされるのでメッセージスキップ
				}
			}

		case cli := <-entering:
			for client := range clients {
				cli.ch <- client.name + " is there!"
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

//!-broadcaster

func getName(conn net.Conn) string {
	input := bufio.NewScanner(conn)
	fmt.Fprint(conn, "enter your name: ")
	if input.Scan() {
		return input.Text()
	}
	return conn.RemoteAddr().String()
}

//!+handleConn
func handleConn(conn net.Conn) {
	ch := make(chan string, 20) // outgoing client messages, メッセージスキップが起こらないように適当にバッファを作成
	go clientWriter(conn, ch)

	client := client{getName(conn), ch}
	client.ch <- "You are " + client.name
	messages <- client.name + " has arrived"
	entering <- client

	input := bufio.NewScanner(conn)
	event := make(chan struct{})
	go func() {
		for {
			select {
			case <-time.After(5 * time.Minute):
				fmt.Fprintln(conn, "You were disconnected due to no action for 5 minutes")
				conn.Close()
				for range event {
					// do nothing
				}
				break
			case <-event:
			}
		}
	}()
	for input.Scan() {
		event <- struct{}{}
		messages <- client.name + ": " + input.Text()
	}
	close(event)
	// NOTE: ignoring potential errors from input.Err()

	leaving <- client
	messages <- client.name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main
