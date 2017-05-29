// Copyright 2017 Ken Mirua
package ex05

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestCountWordsAndImages(t *testing.T) {
	input := `<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Document</title>
		</head>
		<body>
			This is test for words count.
			<img src="image1.png">
			<img src="image2.png">
		</body>
		</html>`
	expectedWordsCount := 7  // Document, This, is, test, for, words, count.
	expectedImagesCount := 2 // <img src="image1.png">, <img src="image2.png">

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Fatalf("parsing html: %v", err)
	}

	actualWordsCount, actualImagesCount := CountWordsAndImages(doc)

	if expectedWordsCount != actualWordsCount {
		t.Fatalf("word count: Expected is %d but actual is %d.", expectedWordsCount, actualWordsCount)
	}

	if expectedImagesCount != actualImagesCount {
		t.Fatalf("image count: Expected is %d but actual is %d.", expectedImagesCount, actualImagesCount)
	}

}
