// Copyright 2017 Ken Miura
// 画面表示にansi escape codeを利用。なのでコマンドプロンプトでは動作しない。PowerShell上で動作させる。
package main

import (
	"fmt"
	"io"
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

func WriteClockWithScreenLock(timeZone string, i int, buf []byte, nr int) (int, error) {
	screenLock.Lock()
	defer screenLock.Unlock()
	fmt.Printf("\x1b[%d;0H", i+1)
	fmt.Printf(timeZone + ": ")
	return os.Stdout.Write(buf[0:nr])
}

func main() {
	clearScreen()
	moveCursorToOrigin()
	var wg sync.WaitGroup
	for i, arg := range os.Args[1:] {
		wg.Add(1)
		go func(i int, arg string) {
			defer wg.Done()
			tzAndHost := strings.Split(arg, "=")
			timeZone := tzAndHost[0]
			host := tzAndHost[1]
			conn, err := net.Dial("tcp", host)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			buf := make([]byte, 1024)
			for {
				nr, er := conn.Read(buf)
				if nr > 0 {
					nw, ew := WriteClockWithScreenLock(timeZone, i, buf, nr)
					if ew != nil {
						err = ew
						break
					}
					if nr != nw {
						err = io.ErrShortWrite
						break
					}
				}
				if er != nil {
					if er != io.EOF {
						err = er
					}
					break
				}
			}
		}(i, arg)
	}
	wg.Wait()
}
