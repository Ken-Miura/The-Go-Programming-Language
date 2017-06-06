// Copyright 2017 Ken Miura
package ex02

import "golang.org/x/net/html"

func CountElement(elementMap map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return elementMap
	}
	if n.Type == html.ElementNode {
		if elementMap == nil {
			elementMap = make(map[string]int)
		}
		elementMap[n.Data]++
	}
	return CountElement(CountElement(elementMap, n.FirstChild), n.NextSibling)
}
