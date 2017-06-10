// Copyright 2017 Ken Miura
package ex01_test

import (
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch06/ex01"
)

func TestNewIntSet(t *testing.T) {
	tests := []struct {
		integers []int
	}{
		{[]int{}},
		{[]int{0}},
		{[]int{1}},
		{[]int{1, 2, 3, 4}},
		{[]int{63}},
		{[]int{64}},
		{[]int{127, 128}},
	}

	for _, test := range tests {
		actual := ex01.NewIntSet(test.integers...)
		if actual.Len() != len(test.integers) {
			t.Fatalf("(*ex01.IntSet).Len() in ex01.NewIntSet failed: expected is %d but actual is %d", len(test.integers), actual.Len())
		}
		for _, integer := range test.integers {
			if !actual.Has(integer) {
				t.Fatalf("(*ex01.IntSet).Has(int) in ex01.NewIntSet failed: actual is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		receiver *ex01.IntSet
		expected int
	}{
		{ex01.NewIntSet(), 0},
		{ex01.NewIntSet(0), 1},
		{ex01.NewIntSet(0, 1), 2},
		{ex01.NewIntSet(64), 1},
		{ex01.NewIntSet(64, 128, 192, 256), 4},
		{ex01.NewIntSet(2, 2), 1},
	}

	for _, test := range tests {
		actual := test.receiver.Len()
		if actual != test.expected {
			t.Fatalf("(*ex01.IntSet).Len() failed: expected is %d but actual is %d", test.expected, actual)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	tests := []struct {
		receiver *ex01.IntSet
		input    int
		expected []int
	}{
		{ex01.NewIntSet(0), 0, []int{}},
		{ex01.NewIntSet(1), 1, []int{}},
		{ex01.NewIntSet(0, 1), 2, []int{0, 1}},
		{ex01.NewIntSet(64), 64, []int{}},
		{ex01.NewIntSet(0, 64, 192), 128, []int{0, 64, 192}},
	}

	for _, test := range tests {
		test.receiver.Remove(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*ex01.IntSet).Len() in (*ex01.IntSet).Remove(int) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for _, integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*ex01.IntSet).Has(int) in *ex01.IntSet).Remove(int) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	tests := []struct {
		receiver *ex01.IntSet
		values   []int
		expected []int
	}{
		{ex01.NewIntSet(), []int{}, []int{}},
		{ex01.NewIntSet(1), []int{1}, []int{}},
		{ex01.NewIntSet(0, 1), []int{0, 1}, []int{}},
		{ex01.NewIntSet(64), []int{64}, []int{}},
		{ex01.NewIntSet(0, 64, 128, 192), []int{0, 64, 128, 192}, []int{}},
	}

	for _, test := range tests {
		test.receiver.Clear()
		if test.receiver.Len() != 0 {
			t.Fatalf("(*ex01.IntSet).Clear() failed: receiver is expected to be length 0 but length %d", test.receiver.Len())
		}
		for _, value := range test.values {
			if test.receiver.Has(value) {
				t.Fatalf("(*ex01.IntSet).Clear() failed: receiver is expected to have no valeu but has %d", value)
			}
		}
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		receiver *ex01.IntSet
		expected []int
	}{
		{ex01.NewIntSet(), []int{}},
		{ex01.NewIntSet(1), []int{1}},
		{ex01.NewIntSet(0, 1), []int{0, 1}},
		{ex01.NewIntSet(64), []int{64}},
		{ex01.NewIntSet(0, 64, 128, 192), []int{0, 64, 128, 192}},
	}

	for _, test := range tests {
		actual := test.receiver.Copy()
		if actual.Len() != len(test.expected) {
			t.Fatalf("(*ex01.IntSet).Len() in (*ex01.IntSet).Copy() failed: expected is %d but actual is %d", len(test.expected), actual.Len())
		}
		for _, integer := range test.expected {
			if !actual.Has(integer) {
				t.Fatalf("(*ex01.IntSet).Has() in (*ex01.IntSet).Copy() failed: actual is expected to have %d", integer)
			}
		}
	}
}
