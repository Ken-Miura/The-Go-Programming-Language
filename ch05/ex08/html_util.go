// Copyright 2017 Ken Miura
package ex08

import (
	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) (found *html.Node) {
	forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == id {
			found = n
			return false
		}
		return true
	}, nil)
	return found
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil && !pre(n) {
		return false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if !forEachNode(c, pre, post) {
			return false
		}
	}

	if post != nil && !post(n) {
		return false
	}
	return true
}
