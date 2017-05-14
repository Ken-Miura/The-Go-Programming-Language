// Copyright 2017 Ken Mirua
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const APIRoot = "http://www.omdbapi.com/?t="

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, `usage: `+os.Args[0]+` "movie title"`)
		fmt.Fprintln(os.Stderr, `ex. `+os.Args[0]+` "The Lord of the Rings: The Fellowship of the Ring"`)
		return
	}

	q := url.QueryEscape(os.Args[1])
	resp1, err := http.Get(APIRoot + q)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp1.Body.Close()

	var info movieInfo
	if err := json.NewDecoder(resp1.Body).Decode(&info); err != nil {
		fmt.Fprintf(os.Stderr, "failed to decode response as json. error: %v\n", err)
		return
	}

	if info.Response == "False" {
		fmt.Fprintf(os.Stderr, "title (%s) not found\n", os.Args[1])
		return
	}

	resp2, err := http.Get(info.Poster)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer resp2.Body.Close()
	io.Copy(os.Stdout, resp2.Body)
}

type movieInfo struct {
	Response string
	Title    string
	Poster   string
}
