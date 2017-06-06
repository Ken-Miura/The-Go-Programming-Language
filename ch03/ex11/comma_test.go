// Copyright 2017 Ken Miura
package ex11

import (
	"testing"
)

var tests = []struct {
	input    string // input
	expected string //  expected
}{
	{"-1", "-1"},
	{"12", "12"},
	{"123", "123"},
	{"-1234", "-1,234"},
	{"1234567890", "1,234,567,890"},
	{"-1234.0", "-1,234.0"},
	{"12345.67890", "12,345.67890"},
}

func TestComma(t *testing.T) {
	for _, tt := range tests {
		actual := Comma(tt.input)
		if actual != tt.expected {
			t.Fatalf("expected is '%s', but actual is '%s'", tt.expected, actual)
		}
	}
}
