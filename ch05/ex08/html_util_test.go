// Copyright 2017 Ken Mirua
package ex08

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestElementByID(t *testing.T) {
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
			<p>third p tag</p>
		</body>
		</html>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parsing html: %v", err)
	}

	actual := ElementByID(doc, "p")

	if actual.Data != "p" {
		t.Fatalf("failed to find element by id: %s", actual.Data)
	}
	if actual.FirstChild.Data != "first p tag" {
		t.Fatalf("failed to suspend operation when ElementByID func found element: %s", actual.FirstChild.Data)
	}
}
