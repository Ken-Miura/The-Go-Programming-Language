// Copyright 2017 Ken Miura
package ex17

import (
	"golang.org/x/net/html"
)

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	var nodes []*html.Node
	for _, s := range name {
		nodes = append(nodes, elementByTagName(doc, s)...)
	}
	return nodes
}

func elementByTagName(doc *html.Node, name string) []*html.Node {
	var nodes []*html.Node
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == name {
			nodes = append(nodes, n)
		}
	}, nil)
	return nodes
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
