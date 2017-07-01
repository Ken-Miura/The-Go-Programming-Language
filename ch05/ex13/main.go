// Copyright 2017 Ken Miura
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

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}

	reg := regexp.MustCompile(`https?://`)
	protocolIndexes := reg.FindIndex([]byte(url))
	protocol := url[:protocolIndexes[1]]
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

func main() {
	for _, url := range os.Args[1:] {
		crawl(url)
	}
}
