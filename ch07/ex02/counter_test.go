// Copyright 2017 Ken Miura
package ex02

import (
	"bytes"
	"testing"
)

var tests = []struct {
	input    string
	expected int64
}{
	{"", 0},
	{"hello", 5},
	{" ", 1},
	{"world", 5},
}

func TestCountingWriter(t *testing.T) {

	for _, test := range tests {
		var buf bytes.Buffer
		writer, pCounter := CountingWriter(&buf)
		_, _ = writer.Write([]byte(test.input))
		if *pCounter != test.expected {
			t.Fatalf("%d is expected but result1 is %d", test.expected, *pCounter)
		}
	}
}
