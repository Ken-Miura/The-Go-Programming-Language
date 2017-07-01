// Copyright 2017 Ken Miura
// crawl2を修正して作成
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

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

	reg := regexp.MustCompile(`https?://`)
	protocolIndexes := reg.FindIndex([]byte(url))
	protocol := ""
	if protocolIndexes != nil {
		protocol = url[:protocolIndexes[1]]
	}
	hostName := ""
	if protocolIndexes != nil {
		hostName = url[protocolIndexes[1]:]
	}
	i := strings.Index(hostName, "/")
	if i != -1 {
		hostName = hostName[:i]
	}

	if err := os.Mkdir(hostName, 0777); err != nil && !os.IsExist(err) {
		log.Print(err)
		return list
	}

	for _, v := range list {
		func() {
			if !strings.HasPrefix(v, protocol+hostName) {
				return
			}
			resp, err := http.Get(v)
			if err != nil {
				log.Print(err)
			}
			defer resp.Body.Close()
			local := path.Base(resp.Request.URL.Path)
			if local == "/" {
				local = "index.html"
			}
			f, err := os.Create(hostName + string(filepath.Separator) + local)
			defer func() {
				if closeErr := f.Close(); err == nil {
					err = closeErr
				}
			}()
			_, err = io.Copy(f, resp.Body)
		}()
	}

	return list
}

//!-sema

//!+
func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

//!-
