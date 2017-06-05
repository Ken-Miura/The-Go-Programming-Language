// Copyright 2017 Ken Mirua
package ex10

import (
	"sort"
	"testing"
)

var tests = []struct {
	input    string
	expected bool
}{
	{"", true},
	{"civic", true},
	{"level", true},
	{"refer", true},
	{"test", false},
	{"function", false},
	{"こんにちは", false},
	{"まさかさかさま", true},
}

type runeSlice []rune

var _ sort.Interface = (runeSlice)(nil)

func (rs runeSlice) Len() int {
	return len(rs)
}

func (rs runeSlice) Less(i, j int) bool {
	return rs[i] < rs[j]
}

func (rs runeSlice) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func TestIsPalindrome(t *testing.T) {
	for _, test := range tests {
		result := IsPalindrome(runeSlice(test.input))
		if result != test.expected {
			t.Errorf("test %s: %t is expected but result is %t", test.input, test.expected, result)
		}
	}
}
