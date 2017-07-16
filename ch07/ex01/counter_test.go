// Copyright 2017 Ken Miura
package ex01

import (
	"testing"
)

func TestWordCounter_Write(t *testing.T) {
	tests := []struct {
		input     string
		expected1 int
		expected2 int
	}{
		{"", 0, 0},
		{"one", 3, 1},
		{"one two", 7, 2},
		{"one two three", 13, 3},
	}

	for _, test := range tests {
		writer := new(WordCounter)
		result, _ := writer.Write([]byte(test.input))
		if result != test.expected1 {
			t.Fatalf("incorrect byte count: %d is expected but result is %d", test.expected1, result)
		}
		if writer.WordCount() != test.expected2 {
			t.Fatalf("incorrect word count: %d is expected but result is %d", test.expected2, writer.WordCount())
		}
	}
}

func TestLineCounter_Write(t *testing.T) {
	tests := []struct {
		input     string
		expected1 int
		expected2 int
	}{
		{``, 0, 0},
		{`one`, 3, 1},
		{`one
			two`, 10, 2},
		{`one
			two
			three`, 19, 3},
	}

	for _, test := range tests {
		writer := new(LineCounter)
		result, _ := writer.Write([]byte(test.input))
		if result != test.expected1 {
			t.Fatalf("incorrect byte count: %d is expected but result is %d", test.expected1, result)
		}
		if writer.LineCount() != test.expected2 {
			t.Fatalf("incorrect line count: %d is expected but result is %d", test.expected2, result)
		}
	}
}
