// Copyright 2017 Ken Miura
package ex07

import "testing"

var tests = []struct {
	input    []byte
	expected []byte
}{
	{[]byte("abc"), []byte("cba")},
	{[]byte("世界"), []byte("界世")},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		actual := Reverse(tt.input)
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
