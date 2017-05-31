// Copyright 2017 Ken Mirua
package ex01

import "testing"

func TestWordCounter_Write(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"one", 1},
		{"one two", 2},
		{"one two three", 3},
	}

	writer := new(WordCounter)
	for _, test := range tests {
		result, _ := writer.Write([]byte(test.input))
		if result != test.expected {
			t.Fatalf("%d is expected but result is %d", test.expected, result)
		}
	}
}

func TestLineCounter_Write(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{``, 0},
		{`one`, 1},
		{`one
			two`, 2},
		{`one
			two
			three`, 3},
	}

	writer := new(LineCounter)
	for _, test := range tests {
		result, _ := writer.Write([]byte(test.input))
		if result != test.expected {
			t.Fatalf("%d is expected but result is %d", test.expected, result)
		}
	}
}
