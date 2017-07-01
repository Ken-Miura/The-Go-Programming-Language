// Copyright 2017 Ken Miura
// crawl2を修正して作成
// TODO 5-13とセットでバグ修正
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
	defer func() {
		<-tokens // release the token
	}()
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	reg := regexp.MustCompile(`https?://`)
	httpIndexes := reg.FindIndex([]byte(url))
	httpScheme := ""
	if httpIndexes != nil {
		httpScheme = url[:httpIndexes[1]]
	}
	hostName := ""
	if httpIndexes != nil {
		hostName = url[httpIndexes[1]:]
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
			if !strings.HasPrefix(v, httpScheme+hostName) {
				return
			}
			resp, err := http.Get(v)
			if err != nil {
				log.Print(err)
			}
			defer resp.Body.Close()
			dir, file := path.Split(resp.Request.URL.Path)
			if dir != "" {
				if err := os.Mkdir(hostName+string(filepath.Separator)+dir, 0777); err != nil && !os.IsExist(err) {
					log.Print(err)
					return
				}
			}
			if file == "" {
				file = "index.html"
			}
			fileName := ""
			if dir != "" {
				fileName = hostName + string(filepath.Separator) + dir + string(filepath.Separator) + file
			} else {
				fileName = hostName + string(filepath.Separator) + file
			}
			f, err := os.Create(fileName)
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
