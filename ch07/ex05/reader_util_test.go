// Copyright 2017 Ken Miura
package ex05

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

var tests = []struct {
	input1    string
	input2    int64
	expected1 string
	expected2 int64
}{
	{"", 1, "", 0},
	{"hello", 4, "hell", 4},
	{"hello", 5, "hello", 5},
	{"hello", 6, "hello", 5},
	{"hello", 0, "", 0},
}

func TestLimitReader(t *testing.T) {

	for _, test := range tests {
		r := LimitReader(strings.NewReader(test.input1), test.input2)
		var out bytes.Buffer
		nbytes, _ := io.Copy(&out, r)
		if out.String() != test.expected1 {
			t.Fatalf("%s is expected but result is %s", test.expected1, out.String())
		}
		if nbytes != test.expected2 {
			t.Fatalf("%d is expected but result is %d", test.expected2, nbytes)
		}
	}
}
