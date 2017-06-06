// Copyright 2017 Ken Miura
package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

func Parse(input string) (*html.Node, error) {
	return html.Parse(NewReader(input))
}

func NewReader(input string) io.Reader {
	return &stringReader{input, 0}
}

type stringReader struct {
	str    string
	offset int
}

var _ io.Reader = (*stringReader)(nil)

func (sr *stringReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		if (i + sr.offset) == len(sr.str) {
			return i, io.EOF
		}
		p[i] = sr.str[i+sr.offset]
	}
	sr.offset += len(p)
	return len(p), nil
}

// 以下サンプルコードfindlinks2よりほぼそのまま引用
// レスポンスをパースする際、io.Readerからではなく、stringから読み込むParseを利用するように修正
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// io.Readerからではなく、stringから読み込むParseを利用するように修正
	//doc, err := html.Parse(resp.Body)
	var buf bytes.Buffer
	io.Copy(&buf, resp.Body)
	doc, err := Parse(buf.String())

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}
