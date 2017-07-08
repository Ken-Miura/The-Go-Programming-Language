// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"os"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "hello"
		for {
			message := <-ch2
			ch1 <- message
			fmt.Println("p1")
		}
	}()

	go func() {
		for {
			message := <-ch1
			ch2 <- message
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
