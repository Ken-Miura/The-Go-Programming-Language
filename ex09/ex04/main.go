// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"time"
)

var n = 3884055

func main() {
	current := make(chan int)
	first := current
	var last chan int
	for i := 0; i < n; i++ {
		next := make(chan int)
		go func(current, next chan int) {
			value := <-current
			next <- value
		}(current, next)
		current = next
		last = current
	}

	start := time.Now()
	first <- 0
	<-last
	fmt.Printf("passing %d channels took %s", n, time.Since(start))
}
