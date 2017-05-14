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
		fmt.Println(`usage: ` + os.Args[0] + ` "movie title"`)
		fmt.Println(`ex. ` + os.Args[0] + ` "The Lord of the Rings: The Fellowship of the Ring"`)
		return
	}

	q := url.QueryEscape(os.Args[1])
	resp1, err := http.Get(APIRoot + q)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp1.Body.Close()

	var info movieInfo
	if err := json.NewDecoder(resp1.Body).Decode(&info); err != nil {
		fmt.Printf("failed to decode response as json. error: %v", err)
		return
	}

	if info.Response == "False" {
		fmt.Println("Not found")
		return
	}

	resp2, err := http.Get(info.Poster)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp2.Body.Close()

	f, err := os.Create("poster.jpg") // 画像ファイルの拡張子の仕様はjpgで固定？？？
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, resp2.Body)
}

type movieInfo struct {
	Response string
	Title    string
	Poster   string
}
