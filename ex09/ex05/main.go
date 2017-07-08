// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	abort := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		times := 0
		for {
			select {
			case message := <-ch2:
				ch1 <- message
				times++
			case <-abort:
				close(ch1)
				for range ch2 {

				}
				fmt.Printf("first goroutine has received and sent %d times for 1s\n", times)
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			times := 0
			for {
				select {
				case message := <-ch1:
					ch2 <- message
					times++
				case <-abort:
					close(ch2)
					for range ch1 {

					}
					fmt.Printf("second goroutine has received and sent %d times for 1s\n", times)
					return
				}
			}
		}
	}()

	ch1 <- "hello"
	select {
	case <-time.After(1 * time.Second):
		close(abort)
	}
	wg.Wait()
}
