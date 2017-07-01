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

func main() {
	for _, url := range os.Args[1:] {
		crawl(url)
	}
}
