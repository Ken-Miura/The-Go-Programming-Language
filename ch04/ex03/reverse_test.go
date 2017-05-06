// Copyright 2017 Ken Mirua
package ex03

import "testing"

var tests = []struct {
	input    [10]int
	expected [10]int
}{
	{[10]int{0: 1}, [10]int{9: 1}},
	{[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
}

func TestReverse(t *testing.T) {
	for _, tt := range tests {
		Reverse(&tt.input)
		if tt.input != tt.expected {
			t.Fatalf("expected is %v, but actual is %v", tt.expected, tt.input)
		}
	}
}
