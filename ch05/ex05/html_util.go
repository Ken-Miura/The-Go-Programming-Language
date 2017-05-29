// Copyright 2017 Ken Mirua
package ex05

import (
	"bufio"
	"golang.org/x/net/html"
	"strings"
)

func CountWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode {
		input := bufio.NewScanner(strings.NewReader(n.Data))
		input.Split(bufio.ScanWords)
		for input.Scan() {
			words++
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		numOfWords, numOfImages := CountWordsAndImages(c)
		words += numOfWords
		images += numOfImages
	}
	return
}
