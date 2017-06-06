// Copyright 2017 Ken Miura
package main

import (
	"bytes"
	"golang.org/x/net/html"
	"testing"
)

var tests = []struct {
	url string
}{
	{"https://stackoverflow.com"},
	{"https://golang.org"},
}

func TestOutline(t *testing.T) {
	for _, test := range tests {
		buf := new(bytes.Buffer)
		outline(buf, test.url)
		_, err := html.Parse(buf)
		if err != nil {
			t.Fatalf("parsing %s as HTML: %v", test.url, err)
		}
	}
}
