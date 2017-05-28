// Copyright 2017 Ken Mirua
package ex02

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestCountElement(t *testing.T) {

	input := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Document</title>
		</head>
		<body>
			<p>first p tag</p>
			<p>second p tag</p>
			<span style="background-color: #0099FF">blue</span>
			<span style="background-color: #FFFF00">yellow</span>
			<span style="background-color: #33CC33">green</span>
			<div align="center">first line<br>second line<br></div>
		</body>
		</html>`
	expected := map[string]int{"html": 1, "head": 1, "body": 1, "p": 2, "span": 3, "div": 1, "br": 2, "meta": 1, "title": 1}

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parsing html: %v", err)
	}

	actual := CountElement(nil, doc)

	if len(expected) != len(actual) {
		t.Fatalf("Number of elements is wrong. Expected is %d but actual is %d", len(expected), len(actual))
	}

	for element, count := range expected {
		actualCount, ok := actual[element]
		if !ok {
			t.Fatalf("element %s is not found in result", element)
		}
		if actualCount != count {
			t.Fatalf("count of element (%s) is wrong. %d is expected but actual is %d", element, count, actualCount)
		}
	}
}
