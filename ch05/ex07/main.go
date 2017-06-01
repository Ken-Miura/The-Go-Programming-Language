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

func main() {
	for _, url := range os.Args[1:] {
		outline(os.Stdout, url)
	}
}

func outline(out io.Writer, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(out, doc, startElement, endElement)

	return nil
}

func forEachNode(out io.Writer, n *html.Node, pre, post func(out io.Writer, n *html.Node)) {
	if pre != nil {
		pre(out, n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(out, c, pre, post)
	}

	if post != nil {
		post(out, n)
	}
}

var depth int

func startElement(out io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild == nil {
			fmt.Fprintf(out, "%*s<%s%s/>\n", depth*2, "", n.Data, attributes(n))
		} else {
			fmt.Fprintf(out, "%*s<%s%s>\n", depth*2, "", n.Data, attributes(n))
		}
		if n.FirstChild != nil && n.FirstChild.Type == html.TextNode {
			textData := strings.TrimSpace(n.FirstChild.Data)
			if textData != "" {
				fmt.Fprintf(out, "%*s%s\n", depth*3, "", textData)
			}
		}
		depth++
	}
}

func attributes(n *html.Node) string {
	s := ""
	for _, attr := range n.Attr {
		s += fmt.Sprintf(" %s", attr.Key)
		if attr.Val != "" {
			s += fmt.Sprintf(`="%s"`, attr.Val)
		}
	}
	return s
}

func endElement(out io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			return
		}
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
