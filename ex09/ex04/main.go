// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"time"
)

// メモリ16GB、OS Windows 10で3883996個のチャネルとそれらをつなぐgo routineをout of memoryを起こすことなく作成できた。それ以上でout of memoryでパニックが発生した。
var n = 3883996

func main() {
	current := make(chan int)
	first := current
	var last chan int
	for i := 0; i < n-1; i++ {
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
	// メモリ16GB、OS Windows 10、CPU i7 6700 3.4GHzで下記の出力となった。
	// 一回目：passing 3883996 channels took 1m18.5663058s
	// 二回目：passing 3883996 channels took 1m14.468393s
	fmt.Printf("passing %d channels took %s", n, time.Since(start))
}
