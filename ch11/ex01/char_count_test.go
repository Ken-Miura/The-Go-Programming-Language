// Copyright 2017 Ken Miura
package ex01_test

import (
	"io"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch11/ex01"
)

var tests = []struct {
	input io.Reader
	want1 map[rune]int
	want2 [utf8.UTFMax + 1]int
	want3 int
	want4 error
}{
	{strings.NewReader(""), make(map[rune]int), [utf8.UTFMax + 1]int{}, 0, nil},
	{strings.NewReader("a"), make(map[rune]int), [utf8.UTFMax + 1]int{}, 0, nil},
}

func TestCharCount(t *testing.T) {
	for _, test := range tests {
		got1, got2, got3, got4 := ex01.CharCount(test.input)

		// counts of Unicode characters
		if len(got1) != len(test.want1) {
			t.Errorf("num of Unicode characters was wrong (got %d, want %d)", len(got1), len(test.want1))
		}
		for k1, v1 := range test.want1 {
			v2, ok := got1[k1]
			if !ok {
				t.Errorf("Unicode character (%q) was not counted", k1)
			}
			if v2 != v1 {
				t.Errorf("counts of Unicode character (%q) were wrong (got %d, want %d)", k1, got1[k1], test.want1[k1])
			}
		}

		// counts of lengths of UTF-8 encodings
		for i := range test.want2 {
			if got2[i] != test.want2[i] {
				t.Errorf("counts of lengths of UTF-8 encodings (%d bytes) were wrong (got %d, want %d)", i, got2[i], test.want2[i])
			}
		}

		// counts of invalid UTF-8 characters
		if got3 != test.want3 {
			t.Errorf("counts of invalid UTF-8 characters were wrong (got %d, want %d)", got3, test.want3)
		}

		// error check
		if got4 != test.want4 {
			t.Errorf("wrong error returned (got %v, want %v)", got4, test.want4)
		}
	}
}
