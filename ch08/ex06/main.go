// Copyright 2017 Ken Miura
// crawl2をもとに作成
package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

//!+sema
// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

//!-sema

var depth = flag.Int("depth", 1, "depth for crawling link")

type linkAndDepth struct {
	link  string
	depth int
}

func createLinks(urls []string, depth int) []linkAndDepth {
	var l []linkAndDepth
	for _, url := range urls {
		l = append(l, linkAndDepth{url, depth})
	}
	return l
}

//!+
func main() {
	flag.Parse()
	if *depth < 0 {
		fmt.Println("depth must be 0 or more")
		return
	}
	worklist := make(chan []linkAndDepth)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- createLinks(flag.Args(), 0)
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if link.depth > *depth {
				continue
			}
			if !seen[link.link] {
				seen[link.link] = true
				n++
				go func(link linkAndDepth) {
					worklist <- createLinks(crawl(link.link), link.depth+1)
				}(link)
			}
		}
	}
}

//!-
