// Copyright 2017 Ken Miura
package ex04

import "testing"

var tests = []struct {
	input1   []int
	input2   int
	expected []int
}{
	{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
	{[]int{1, 2, 3}, -1, []int{2, 3, 1}},
	{[]int{1, 2, 3}, 0, []int{1, 2, 3}},
}

func TestRotate(t *testing.T) {
	for _, tt := range tests {
		actual := Rotate(tt.input1, tt.input2)
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
