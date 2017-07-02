// Copyright 2017 Ken Miura
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 複数のURLのリクエストを送り、もっともはやいレスポンスのものを出力する
func main() {
	if len(os.Args) == 1 {
		fmt.Println("usage: " + os.Args[0] + " 'urls'")
		fmt.Println("ex: " + os.Args[0] + " https://golang.org https://www.stackoverflow.com")
		return
	}
	fmt.Println(fetch(os.Args[1:]))
}

func fetch(list []string) string {
	responses := make(chan string, len(list))
	for _, link := range list {
		go func(link string) {
			resp, err := http.Get(link)
			if err != nil {
				log.Printf("fetching %s: %v", link, err)
				return
			}
			defer resp.Body.Close()
			var buf bytes.Buffer
			n, err := io.Copy(&buf, resp.Body)
			if err != nil {
				log.Printf("reading from %s at %d bytes: %v", link, n, err)
				return
			}
			responses <- buf.String()
		}(link)
	}
	return <-responses
}
