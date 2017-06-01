// Copyright 2017 Ken Mirua
package main

import "testing"

func TestMin(t *testing.T) {
	tests := []struct {
		firstVal int
		values   []int
		expected int
	}{
		{1, []int{2, 3, 4, 5}, 1},
		{3, []int{2, 3, 6, 10}, 2},
		{1, []int{2, 3, -4, 5}, -4},
		{3, []int{2, 0, 6, 10}, 0},
	}

	for _, test := range tests {
		result := Min(test.firstVal, test.values...)
		if result != test.expected {
			t.Fatalf("%d is expected but result is %d", test.expected, result)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		firstVal int
		values   []int
		expected int
	}{
		{1, []int{2, 3, 4, 5}, 5},
		{3, []int{2, 3, 6, 10}, 10},
		{1, []int{2, 3, -4, 5}, 5},
		{3, []int{2, 0, 6, 10}, 10},
	}

	for _, test := range tests {
		result := Max(test.firstVal, test.values...)
		if result != test.expected {
			t.Fatalf("%d is expected but result is %d", test.expected, result)
		}
	}
}
