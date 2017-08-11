// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"time"

	"sync"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch09/ex03/memo"
)

func main() {
	pMemo := memo.New(test)
	defer pMemo.Close()

	// キャッシュなしの処理速度計測
	cancel := make(chan struct{})
	done := make(chan struct{})
	var wg1 sync.WaitGroup
	wg1.Add(1)
	go func() {
		defer wg1.Done()
		<-done
		pMemo.Get("key1", cancel)
	}()
	close(cancel)
	done <- struct{}{}
	wg1.Wait()
	// この時点でkey1を使ったGet呼び出しはキャンセルされて終了。なのでキャッシュを持っていない。
	start := time.Now()
	pMemo.Get("key1", make(chan struct{}))
	fmt.Printf("Get process time (no cache): %s\n", time.Since(start))

	// キャッシュ有りの処理速度計測
	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		pMemo.Get("key2", nil)
	}()
	wg2.Wait() // Get呼び出しが完了し、key2のキャッシュができるように待機
	start = time.Now()
	pMemo.Get("key2", nil)
	fmt.Printf("Get process time (has cache): %s\n", time.Since(start))
}

func test(_ string, _ <-chan struct{}) (interface{}, error) {
	time.Sleep(1 * time.Second)
	return nil, nil
}
