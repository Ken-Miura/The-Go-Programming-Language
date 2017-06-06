// Copyright 2017 Ken Miura
// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		bytesCopied, err := io.Copy(os.Stdout, resp.Body)
		fmt.Fprintf(os.Stderr, "status code: "+resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v at %d byte(s) copied\n", url, err, bytesCopied)
			os.Exit(1)
		}
	}
}
