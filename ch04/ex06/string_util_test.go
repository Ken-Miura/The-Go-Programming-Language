// Copyright 2017 Ken Miura
package ex06

import "testing"

var tests = []struct {
	input    []byte
	expected []byte
}{
	{[]byte("a b c"), []byte("a b c")},
	{[]byte("a  b  c"), []byte("a b c")},
	{[]byte("a   b   c"), []byte("a b c")},
	{[]byte("  "), []byte(" ")},
	{[]byte("a　b"), []byte("a b")},
	{[]byte("a\u200A b"), []byte("a b")},
	{[]byte("abc"), []byte("abc")},
	{[]byte("　 abc   　"), []byte(" abc ")},
}

func TestCompressSpaces(t *testing.T) {
	for _, tt := range tests {
		actual := CompressSpaces(tt.input)
		if len(actual) != len(tt.expected) {
			t.Fatalf("expected is %v, but actual is %v", tt.expected, actual)
		}
		for i := range tt.expected {
			if actual[i] != tt.expected[i] {
				t.Fatalf("expected is %v, but actual is %v", tt.expected, actual)
			}
		}
	}
}
