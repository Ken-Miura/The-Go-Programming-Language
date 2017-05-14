// Copyright 2017 Ken Mirua
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
	'search' searches comic including specified word in database`)

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
			fmt.Println("Index file to search comic might not exist.")
			fmt.Println("Please create index file at first.")
			fmt.Println("how to create: " + programName + " -o create")
			return
		}
		var indexes []Index
		ok := constructIndex(f, &indexes)
		if !ok {
			fmt.Println("failed to construct index")
			fmt.Println("Please re-create index file.")
			fmt.Println("how to create: " + programName + " -o create")
			return
		}
		search(indexes, args[0])
	} else {
		panic("This line must not be reached.\n")
	}
}

func createIndex() {
	f, err := os.Create("index.json")
	if err != nil {
		fmt.Printf("failed to create file. error: %v\n", err)
		return
	}
	defer f.Close()

	var indexes []Index
	var ok bool = true
	for i := 1; ok; i++ {
		if i == 404 { // this numbering results in Not Found
			continue
		}
		var index Index
		index, ok = requestIndex(fmt.Sprintf(xkcdAPIRoot, i))
		if ok {
			indexes = append(indexes, index)
		}
	}

	if err := json.NewEncoder(f).Encode(indexes); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("succeeded in creating index")
}

func requestIndex(url string) (Index, bool) {
	fmt.Printf("fetching from: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return Index{}, false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Index{}, false
	}

	var index Index
	if err := json.NewDecoder(resp.Body).Decode(&index); err != nil {
		fmt.Println(err)
		return Index{}, false
	}
	return index, true
}

func constructIndex(r io.Reader, indexes *[]Index) bool {
	if err := json.NewDecoder(r).Decode(&indexes); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func search(indexes []Index, word string) {
	for _, index := range indexes {
		if strings.Contains(index.Title, word) || strings.Contains(index.Transcript, word) {
			fmt.Printf("title: %s\n", index.Title)
			fmt.Printf("URL: %s\n", fmt.Sprintf(xkcdURL, index.Num))
			fmt.Printf("transcript:\n%s\n", index.Transcript)
			return
		}
	}
	fmt.Printf("Not found %s in index\n", word)
}

type Index struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}
