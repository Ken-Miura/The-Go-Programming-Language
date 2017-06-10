// Copyright 2017 Ken Miura
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const xkcdURL = "https://xkcd.com/%d"
const xkcdAPIRoot = xkcdURL + "/info.0.json"

var operation = flag.String("o", "search",
	`operations for xkcd
	'create' creates index by fetching data from "https://xkcd.com/'number'/info.0.json"
	'search' searches comic including specified word`)

func main() {
	flag.Parse()
	if !(*operation == "create" || *operation == "search") {
		fmt.Println("Operation must be create or search.")
		return
	}

	programName := os.Args[0]
	args := flag.Args()
	if *operation == "create" {
		createIndex()
	} else if *operation == "search" {
		if len(args) != 1 {
			fmt.Println("usage: " + programName + " [-o search] 'word'")
			fmt.Println("ex1. " + programName + " Barrel")
			fmt.Println("ex2. " + programName + " -o search Barrel")
			return
		}

		f, err := os.Open("index.json")
		if err != nil {
			fmt.Println(err)
			fmt.Println("index file to search comic might not exist.")
			fmt.Println("Please create index file at first.")
			fmt.Println("how to create: " + programName + " -o create")
			return
		}
		defer f.Close()

		var index []ComicInfo
		ok := constructIndex(f, &index)
		if !ok {
			fmt.Println("failed to construct index")
			fmt.Println("Please re-create index file.")
			fmt.Println("how to create: " + programName + " -o create")
			return
		}
		search(index, args[0])
	} else {
		panic("This line must not be reached.\n")
	}
}

func createIndex() {
	f, err := os.Create("index.json")
	if err != nil {
		fmt.Printf("failed to create file: %v\n", err)
		return
	}
	defer f.Close()

	var index []ComicInfo
	var ok bool = true
	for i := 1; ok; i++ {
		if i == 404 { // this numbering results in Not Found
			continue
		}
		var comicInfo ComicInfo
		comicInfo, ok = requestComicInfo(fmt.Sprintf(xkcdAPIRoot, i))
		if ok {
			index = append(index, comicInfo)
		}
	}

	if err := json.NewEncoder(f).Encode(index); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("succeeded in creating index")
}

func requestComicInfo(url string) (ComicInfo, bool) {
	fmt.Printf("fetching from: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ComicInfo{}, false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ComicInfo{}, false
	}

	var comicInfo ComicInfo
	if err := json.NewDecoder(resp.Body).Decode(&comicInfo); err != nil {
		fmt.Println(err)
		return ComicInfo{}, false
	}
	return comicInfo, true
}

func constructIndex(r io.Reader, indexes *[]ComicInfo) bool {
	if err := json.NewDecoder(r).Decode(&indexes); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func search(index []ComicInfo, word string) {
	for _, comicInfo := range index {
		if strings.Contains(comicInfo.Title, word) || strings.Contains(comicInfo.Transcript, word) {
			fmt.Printf("title: %s\n", comicInfo.Title)
			fmt.Printf("URL: %s\n", fmt.Sprintf(xkcdURL, comicInfo.Num))
			fmt.Printf("transcript:\n%s\n", comicInfo.Transcript)
			return
		}
	}
	fmt.Printf("Not found %s in index\n", word)
}

type ComicInfo struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}
