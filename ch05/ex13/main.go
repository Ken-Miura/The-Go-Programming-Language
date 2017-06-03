// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	dirName := url
	reg := regexp.MustCompile(`https?://`)
	protocolIndexes := reg.FindIndex([]byte(url))
	if protocolIndexes != nil {
		dirName = dirName[protocolIndexes[1]:]
	}
	i := strings.Index(dirName, "/")
	if i != -1 {
		dirName = dirName[:i]
	}

	if err := os.Mkdir(dirName, 0777); err != nil && !os.IsExist(err) {
		log.Print(err)
		return list
	}

	for _, v := range list {
		func() {
			if !strings.HasPrefix(v, url) {
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
			f, err := os.Create(dirName + string(filepath.Separator) + local)
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

func main() {
	for _, url := range os.Args[1:] {
		crawl(url)
	}
}
