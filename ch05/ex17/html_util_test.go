// Copyright 2017 Ken Mirua
package ex17

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestElementByTagName(t *testing.T) {
	input := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Document</title>
		</head>
		<body>
			<p>first p tag</p>
			<span style="background-color: #0099FF">blue</span>
			<img src="image2.png"/>
		</body>
		</html>`

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parsing html: %v", err)
	}

	tags := ElementByTagName(doc, "title", "p", "img")
	if len(tags) != 3 {
		t.Fatalf("%d tags are detected", len(tags))
	}
	titleCount := 0
	pCount := 0
	imgCount := 0
	for _, tag := range tags {
		if tag.Data == "title" {
			if tag.FirstChild.Data != "Document" {
				t.Fatalf("title is 'Document' but %s is detected", tag.FirstChild.Data)
			}
			titleCount++
			if titleCount > 1 {
				t.Fatalf("title tag is detected twice")
			}
		} else if tag.Data == "p" {
			if tag.FirstChild.Data != "first p tag" {
				t.Fatalf("p is 'first p tag' but %s is detected", tag.FirstChild.Data)
			}
			pCount++
			if pCount > 1 {
				t.Fatalf("p tag is detected twice")
			}
		} else if tag.Data == "img" {
			if tag.Attr[0].Val != "image2.png" {
				t.Fatalf("attr src in img is 'image2.png' but %s is detected", tag.FirstChild.Data)
			}
			imgCount++
			if imgCount > 1 {
				t.Fatalf("img tag is detected twice")
			}
		} else {
			t.Fatalf("tag shoud not be detected: %s", tag.Data)
		}
	}
}
