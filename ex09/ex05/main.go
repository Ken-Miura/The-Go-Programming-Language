// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"os"
)

var ch = make(chan string)

func main() {

	go func() {
		ch <- "hello"
		for {
			message := <-ch
			ch <- message
			fmt.Println("p1")
		}
	}()

	go func() {
		for {
			message := <-ch
			ch <- message
			fmt.Println("p2")
		}
	}()

	done := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	<-done
}
