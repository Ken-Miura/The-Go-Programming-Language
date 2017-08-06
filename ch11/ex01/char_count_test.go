// Copyright 2017 Ken Miura
package ex01_test

import (
	"io"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch11/ex01"
)

type testCase struct {
	input io.Reader
	want1 map[rune]int
	want2 [utf8.UTFMax + 1]int
	want3 int
	want4 error
}

var tests = []testCase{
	{strings.NewReader(""), make(map[rune]int), [utf8.UTFMax + 1]int{}, 0, nil},
}

func init() {
	createTestCase2()
	createTestCase3()
	createTestCase4()
}

func createTestCase2() {
	r := strings.NewReader("hello world")

	counts := make(map[rune]int)
	counts['h'] = 1
	counts['e'] = 1
	counts['l'] = 3
	counts['o'] = 2
	counts[' '] = 1
	counts['w'] = 1
	counts['r'] = 1
	counts['d'] = 1

	countsOfEnc := [utf8.UTFMax + 1]int{0, 11, 0, 0, 0}

	countInvalidChar := 0

	var err error = nil

	tests = append(tests, testCase{r, counts, countsOfEnc, countInvalidChar, err})
}

func createTestCase3() {
	r := strings.NewReader("こんにちは、世界")

	counts := make(map[rune]int)
	counts['こ'] = 1
	counts['ん'] = 1
	counts['に'] = 1
	counts['ち'] = 1
	counts['は'] = 1
	counts['、'] = 1
	counts['世'] = 1
	counts['界'] = 1

	countsOfEnc := [utf8.UTFMax + 1]int{0, 0, 0, 8, 0}

	countInvalidChar := 0

	var err error = nil

	tests = append(tests, testCase{r, counts, countsOfEnc, countInvalidChar, err})
}

func createTestCase4() {
	r := strings.NewReader(string([]byte{0xC0})) // 1バイト文字なのに0x7f以上の値を持つもの

	counts := make(map[rune]int)

	countsOfEnc := [utf8.UTFMax + 1]int{0, 0, 0, 0, 0}

	countInvalidChar := 1

	var err error = nil

	tests = append(tests, testCase{r, counts, countsOfEnc, countInvalidChar, err})
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
				t.Errorf("counts of %d bytes UTF-8 encodings were wrong (got %d, want %d)", i, got2[i], test.want2[i])
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
