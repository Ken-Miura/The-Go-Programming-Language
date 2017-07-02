// Copyright 2017 Ken Miura
package main

import (
	"context"
	"fmt"
	"io/ioutil"
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

var done = make(chan struct{})

func fetch(list []string) string {
	defer close(done)
	responses := make(chan string, len(list))
	for _, link := range list {
		go func(link string) {
			req, err := http.NewRequest("GET", link, nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "creating request for %s: %v", link, err)
				return
			}
			ctx := req.Context()
			ctx, cancel := context.WithCancel(ctx)
			req.WithContext(ctx)
			go func() {
				<-done
				cancel()
			}()
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetching %s: %v", link, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				fmt.Fprintf(os.Stderr, "status from %s: %v", link, err)
				return
			}

			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "reading from response (%s): %v", link, err)
				return
			}
			responses <- string(b)
		}(link)
	}
	return <-responses
}
