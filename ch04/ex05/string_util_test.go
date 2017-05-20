// Copyright 2017 Ken Mirua
package ex05

import "testing"

var tests = []struct {
	input    []string
	expected []string
}{
	{[]string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "b", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "b", "c", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "b", "b", "c", "b"}, []string{"a", "b", "c", "b"}},
	{[]string{"a", "b", "b", "c", "b", "b"}, []string{"a", "b", "c", "b"}},
	{[]string{"a", "a", "a"}, []string{"a"}},
}

func TestRemoveAdjacentDuplication(t *testing.T) {
	for _, tt := range tests {
		actual := RemoveAdjacentDuplication(tt.input)
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
