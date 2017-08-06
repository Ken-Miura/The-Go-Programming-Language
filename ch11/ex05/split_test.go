// Copyright 2017 Ken Miura
package ex05

import (
	"strings"
	"testing"
)

var tests = []struct {
	s    string
	sep  string
	want int
}{
	{"", ",", 1},
	{"a", ",", 1},
	{"a:b:c", ":", 3},
	{"a,b,c", ",", 3},
	{"a, b, c", ",", 3},
	{" a,b,c ", ",", 3},
}

func TestSplit(t *testing.T) {
	for _, test := range tests {
		words := strings.Split(test.s, test.sep)
		if len(words) != test.want {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, len(words), test.want)
		}
	}
}
