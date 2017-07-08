// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"time"

	"sync"

	"github.com/Ken-Miura/The-Go-Programming-Language/ex09/ex03/memo"
)

func main() {
	pMemo := memo.New(test)
	defer pMemo.Close()

	done := make(chan struct{})
	canceled := make(chan struct{})
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		<-canceled
		pMemo.Get("key1", done)
	}()
	close(done)
	canceled <- struct{}{}
	wg1.Wait()
	// この時点でkey1を使ったGet呼び出しはキャンセルされて終了。なのでキャッシュを持っていない。

	start := time.Now()
	pMemo.Get("key1", make(chan struct{}))
	fmt.Printf("Get processing time (canceled and has no cache): %s\n", time.Since(start))

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		pMemo.Get("key2", nil)
	}()
	wg2.Wait() // Get呼び出しが完了し、key2のキャッシュができるように待機
	start = time.Now()
	pMemo.Get("key2", nil)
	fmt.Printf("Get processing time (has cache): %s\n", time.Since(start))
}

func test(_ string) (interface{}, error) {
	time.Sleep(1 * time.Second)
	return nil, nil
}
