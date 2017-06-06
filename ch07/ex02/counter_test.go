// Copyright 2017 Ken Miura
package ex02

import (
	"bytes"
	"testing"
)

var tests = []struct {
	input1    string
	expected1 int64
	input2    string
	expected2 int64
	input3    string
	expected3 int64
	input4    string
	expected4 int64
}{
	{"", 0, "hello", 5, " ", 6, "world", 11},
}

func TestCountingWriter(t *testing.T) {

	for _, test := range tests {
		var buf bytes.Buffer
		writer, pCounter := CountingWriter(&buf)
		_, _ = writer.Write([]byte(test.input1))
		if *pCounter != test.expected1 {
			t.Fatalf("%d is expected but result1 is %d", test.expected1, *pCounter)
		}
		_, _ = writer.Write([]byte(test.input2))
		if *pCounter != test.expected2 {
			t.Fatalf("%d is expected but result1 is %d", test.expected2, *pCounter)
		}
		_, _ = writer.Write([]byte(test.input3))
		if *pCounter != test.expected3 {
			t.Fatalf("%d is expected but result1 is %d", test.expected3, *pCounter)
		}
		_, _ = writer.Write([]byte(test.input4))
		if *pCounter != test.expected4 {
			t.Fatalf("%d is expected but result1 is %d", test.expected4, *pCounter)
		}
	}
}
