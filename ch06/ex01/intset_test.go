// Copyright 2017 Ken Mirua
package ex01

import "testing"

func TestIntSet_Len(t *testing.T) {
	tests := []struct {
		receiver IntSet
		expected int
	}{
		{IntSet{[]uint64{0}}, 0},
		{IntSet{[]uint64{1}}, 1},
		{IntSet{[]uint64{3}}, 2},
		{IntSet{[]uint64{0, 1}}, 1},
		{IntSet{[]uint64{1, 1, 1, 1}}, 4},
		{IntSet{[]uint64{4, 0, 5, 3}}, 5},
	}

	for _, test := range tests {
		actual := test.receiver.Len()
		if actual != test.expected {
			t.Fatalf("(*IntSet).Len() failed: expected is %d but actual is %d", test.expected, actual)
		}
	}
}

func TestIntSet_Remove(t *testing.T) {
	tests := []struct {
		receiver IntSet
		input    int
		expected IntSet
	}{
		{IntSet{[]uint64{0}}, 0, IntSet{[]uint64{0}}},
		{IntSet{[]uint64{2}}, 1, IntSet{[]uint64{0}}},
		{IntSet{[]uint64{3}}, 2, IntSet{[]uint64{3}}},
		{IntSet{[]uint64{0: 0, 1: 1}}, 64, IntSet{[]uint64{0: 0, 1: 0}}},
		{IntSet{[]uint64{0: 1, 1: 1, 2: 0, 3: 1}}, 128, IntSet{[]uint64{0: 1, 1: 1, 2: 0, 3: 1}}},
	}

	for _, test := range tests {
		test.receiver.Remove(test.input)
		if len(test.receiver.words) != len(test.expected.words) {
			t.Fatalf("(*IntSet).Remove(int) failed: expected is %v but actual is %v", test.expected, test.receiver)
		}
		for i := range test.receiver.words {
			if test.receiver.words[i] != test.expected.words[i] {
				t.Fatalf("(*IntSet).Remove(int) failed: expected is %v but actual is %v", test.expected, test.receiver)
			}
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	tests := []struct {
		receiver IntSet
		expected IntSet
	}{
		{IntSet{[]uint64{0}}, IntSet{[]uint64{0}}},
		{IntSet{[]uint64{2}}, IntSet{[]uint64{0}}},
		{IntSet{[]uint64{3}}, IntSet{[]uint64{0}}},
		{IntSet{[]uint64{0, 1}}, IntSet{[]uint64{0, 0}}},
		{IntSet{[]uint64{1, 1, 1, 1}}, IntSet{[]uint64{0, 0, 0, 0}}},
	}

	for _, test := range tests {
		test.receiver.Clear()
		if len(test.receiver.words) != len(test.expected.words) {
			t.Fatalf("(*IntSet).Clear() failed: expected is %v but actual is %v", test.expected, test.receiver)
		}
		for i := range test.receiver.words {
			if test.receiver.words[i] != test.expected.words[i] {
				t.Fatalf("(*IntSet).Clear() failed: expected is %v but actual is %v", test.expected, test.receiver)
			}
		}
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		receiver IntSet
		expected IntSet
	}{
		{IntSet{[]uint64{0}}, IntSet{[]uint64{0}}},
		{IntSet{[]uint64{2}}, IntSet{[]uint64{2}}},
		{IntSet{[]uint64{3}}, IntSet{[]uint64{3}}},
		{IntSet{[]uint64{0, 1}}, IntSet{[]uint64{0, 1}}},
		{IntSet{[]uint64{1, 1, 1, 1}}, IntSet{[]uint64{1, 1, 1, 1}}},
	}

	for _, test := range tests {
		actual := test.receiver.Copy()
		if len(actual.words) != len(test.expected.words) {
			t.Fatalf("(*IntSet).Copy() failed: expected is %v but actual is %v", test.expected, actual)
		}
		for i := range actual.words {
			if actual.words[i] != test.expected.words[i] {
				t.Fatalf("(*IntSet).Copy() failed: expected is %v but actual is %v", test.expected, actual)
			}
		}
	}
}
