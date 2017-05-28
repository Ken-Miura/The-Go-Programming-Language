// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"os"
)

// TODO 時間があれば目視確認テストでなくテストコードで
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Printf("parsing html: %v", err)
		os.Exit(1)
	}
	DisplayText(os.Stdout, doc)
}

func DisplayText(out io.Writer, n *html.Node) {
	if n == nil {
		return
	}

	if checkIfNodeShouldBeDisplayed(n) {
		out.Write([]byte(n.Data))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		DisplayText(out, c)
	}
}

func checkIfNodeShouldBeDisplayed(n *html.Node) bool {
	if n.Type != html.TextNode {
		return false
	}
	if n.Parent.Type == html.ElementNode && (n.Parent.Data == "script" || n.Parent.Data == "style") {
		return false
	}
	return true
}
