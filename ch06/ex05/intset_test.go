// Copyright 2017 Ken Miura
package ex05_test

import (
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch06/ex05"
)

func TestIntSet_Elems(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		expected []int
	}{
		{ex05.NewIntSet(), []int{}},
		{ex05.NewIntSet(0), []int{0}},
		{ex05.NewIntSet(1), []int{1}},
		{ex05.NewIntSet(0, 1), []int{0, 1}},
		{ex05.NewIntSet(0, 2, 4), []int{0, 2, 4}},
		{ex05.NewIntSet(0, 1, 2, 3, 4, 5), []int{0, 1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		actual := test.receiver.Elems()
		if len(actual) != len(test.expected) {
			t.Fatalf("(*IntSet).Elems() failed: expected is %v but actual is %v", test.expected, actual)
		}
		for i := range test.expected {
			if actual[i] != test.expected[i] {
				t.Fatalf("(*IntSet).Elems() failed: expected is %v but actual is %v", test.expected, actual)
			}
		}
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		input    *ex05.IntSet
		expected []int
	}{
		{ex05.NewIntSet(), ex05.NewIntSet(), []int{}},
		{ex05.NewIntSet(), ex05.NewIntSet(64), []int{}},
		{ex05.NewIntSet(0), ex05.NewIntSet(0), []int{0}},
		{ex05.NewIntSet(0, 1), ex05.NewIntSet(0, 2), []int{0}},
	}

	for _, test := range tests {
		test.receiver.IntersectWith(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*IntSet).Len() in (*IntSet).IntersectWith(*IntSet) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for _, integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*IntSet).Has() in (*IntSet).IntersectWith(*IntSet) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_DifferenceWith(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		input    *ex05.IntSet
		expected []int
	}{
		{ex05.NewIntSet(), ex05.NewIntSet(), []int{}},
		{ex05.NewIntSet(), ex05.NewIntSet(64), []int{}},
		{ex05.NewIntSet(0), ex05.NewIntSet(1), []int{0}},
		{ex05.NewIntSet(2, 66), ex05.NewIntSet(0, 2), []int{66}},
	}

	for _, test := range tests {
		test.receiver.DifferenceWith(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*IntSet).Len() in (*IntSet).DifferenceWith(*IntSet) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for _, integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*IntSet).Has() in (*IntSet).DifferenceWith(*IntSet) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_SymmetricDifference(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		input    *ex05.IntSet
		expected []int
	}{
		{ex05.NewIntSet(), ex05.NewIntSet(), []int{}},
		{ex05.NewIntSet(), ex05.NewIntSet(64), []int{64}},
		{ex05.NewIntSet(0), ex05.NewIntSet(0), []int{}},
		{ex05.NewIntSet(1, 2), ex05.NewIntSet(1, 4), []int{2, 4}},
	}

	for _, test := range tests {
		test.receiver.SymmetricDifference(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*IntSet).Len() in (*IntSet).SymmetricDifference(*IntSet) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for _, integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*IntSet).Has() in (*IntSet).SymmetricDifference(*IntSet) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_AddAll(t *testing.T) {

	tests := []struct {
		receiver *ex05.IntSet
		input    []int
		expected []int
	}{
		{ex05.NewIntSet(), []int{1, 2, 3}, []int{1, 2, 3}},
		{ex05.NewIntSet(), []int{}, []int{}},
		{ex05.NewIntSet(), []int{64}, []int{64}},
		{ex05.NewIntSet(0), []int{64}, []int{0, 64}},
	}

	for _, test := range tests {
		test.receiver.AddAll(test.input...)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*ex05.IntSet).Len() in (*ex05.IntSet).AddAll(...int) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for _, integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*ex05.IntSet).Has(int) in *ex05.IntSet).AddAll(...int) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

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
		actual := ex05.NewIntSet(test.integers...)
		if actual.Len() != len(test.integers) {
			t.Fatalf("(*ex05.IntSet).Len() in ex05.NewIntSet failed: expected is %d but actual is %d", len(test.integers), actual.Len())
		}
		for _, integer := range test.integers {
			if !actual.Has(integer) {
				t.Fatalf("(*ex05.IntSet).Has(int) in ex05.NewIntSet failed: actual is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		expected int
	}{
		{ex05.NewIntSet(), 0},
		{ex05.NewIntSet(0), 1},
		{ex05.NewIntSet(0, 1), 2},
		{ex05.NewIntSet(64), 1},
		{ex05.NewIntSet(64, 128, 192, 256), 4},
		{ex05.NewIntSet(2, 2), 1},
	}

	for _, test := range tests {
		actual := test.receiver.Len()
		if actual != test.expected {
			t.Fatalf("(*ex05.IntSet).Len() failed: expected is %d but actual is %d", test.expected, actual)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		input    int
		expected []int
	}{
		{ex05.NewIntSet(0), 0, []int{}},
		{ex05.NewIntSet(1), 1, []int{}},
		{ex05.NewIntSet(0, 1), 2, []int{0, 1}},
		{ex05.NewIntSet(64), 64, []int{}},
		{ex05.NewIntSet(0, 64, 192), 128, []int{0, 64, 192}},
	}

	for _, test := range tests {
		test.receiver.Remove(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*ex05.IntSet).Len() in (*ex05.IntSet).Remove(int) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for _, integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*ex05.IntSet).Has(int) in *ex05.IntSet).Remove(int) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		values   []int
		expected []int
	}{
		{ex05.NewIntSet(), []int{}, []int{}},
		{ex05.NewIntSet(1), []int{1}, []int{}},
		{ex05.NewIntSet(0, 1), []int{0, 1}, []int{}},
		{ex05.NewIntSet(64), []int{64}, []int{}},
		{ex05.NewIntSet(0, 64, 128, 192), []int{0, 64, 128, 192}, []int{}},
	}

	for _, test := range tests {
		test.receiver.Clear()
		if test.receiver.Len() != 0 {
			t.Fatalf("(*ex05.IntSet).Clear() failed: receiver is expected to be length 0 but length %d", test.receiver.Len())
		}
		for _, value := range test.values {
			if test.receiver.Has(value) {
				t.Fatalf("(*ex05.IntSet).Clear() failed: receiver is expected to have no valeu but has %d", value)
			}
		}
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		expected []int
	}{
		{ex05.NewIntSet(), []int{}},
		{ex05.NewIntSet(1), []int{1}},
		{ex05.NewIntSet(0, 1), []int{0, 1}},
		{ex05.NewIntSet(64), []int{64}},
		{ex05.NewIntSet(0, 64, 128, 192), []int{0, 64, 128, 192}},
	}

	for _, test := range tests {
		actual := test.receiver.Copy()
		if actual.Len() != len(test.expected) {
			t.Fatalf("(*ex05.IntSet).Len() in (*ex05.IntSet).Copy() failed: expected is %d but actual is %d", len(test.expected), actual.Len())
		}
		for _, integer := range test.expected {
			if !actual.Has(integer) {
				t.Fatalf("(*ex05.IntSet).Has() in (*ex05.IntSet).Copy() failed: actual is expected to have %d", integer)
			}
		}
	}
}
