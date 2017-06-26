// Copyright 2017 Ken Miura
// 画面表示にansi escape codeを利用。なのでコマンドプロンプトでは動作しない。PowerShell上で動作させる。
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func clearScreen() {
	fmt.Printf("\x1b[%dJ", 2)
}

func moveCursorToOrigin() {
	fmt.Printf("\x1b[%d;%dH", 0, 0)
}

var screenLock sync.Mutex

func displayWithScreenLock(str string, x, y int) {
	screenLock.Lock()
	defer screenLock.Unlock()
	fmt.Printf("\x1b[%d;%dH", y, x)
	fmt.Printf(str)
}

func main() {
	clearScreen()
	moveCursorToOrigin()
	var wg sync.WaitGroup
	for i, arg := range os.Args[1:] {
		wg.Add(1)
		go func(i int, arg string) {
			defer wg.Done()
			timeZoneAndHost := strings.Split(arg, "=")
			timeZone := timeZoneAndHost[0]
			host := timeZoneAndHost[1]
			conn, err := net.Dial("tcp", host)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			buf := make([]byte, 1024*4)
			for {
				nbytes, err := conn.Read(buf)
				if nbytes > 0 {
					clockString := string(buf[:nbytes])
					displayWithScreenLock(timeZone+": "+clockString, 0, i+1)
				}
				if err != nil {
					break
				}
			}
		}(i, arg)
	}
	wg.Wait()
}
