// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
	"io"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	outline("https://golang.org")
	//for _, url := range os.Args[1:] {
	//	outline(url)
	//}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(out, "%*s<%s>\n", depth*2, "", n.Data)
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			textData := strings.TrimSpace(n.FirstChild.Data)
			if textData != "" {
				fmt.Fprintf(out, "%*s%s\n", depth*3, "", textData)
			}
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			return
		}
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
