// Copyright 2017 Ken Miura
package ex12

import (
	"testing"
)

var tests = []struct {
	input1   string // input
	input2   string // input
	expected bool   //  expected
}{
	{"ab", "cd", false},
	{"canoe", "ocean", true},
	{"こんにちは", "世界", false},
	{"かとうあい", "あとうかい", true},
}

func TestComma(t *testing.T) {
	for _, tt := range tests {
		actual := IsAnagram(tt.input1, tt.input2)
		if actual != tt.expected {
			t.Fatalf("expected is %t, but actual is %t (input1: %s, input2: %s)", tt.expected, actual, tt.input1, tt.input2)
		}
	}
}
