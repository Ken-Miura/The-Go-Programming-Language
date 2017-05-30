// Copyright 2017 Ken Mirua
package ex01

import "testing"

func TestIntSet_AddAll(t *testing.T) {

	receiver1 := IntSet{[]uint64{0}}
	expected1 := IntSet{[]uint64{14}}

	receiver1.AddAll(1, 2, 3)

	if len(receiver1.words) != len(expected1.words) {
		t.Fatalf("(*IntSet).AddAll() failed: expected is %v but actual is %v", expected1, receiver1)
	}
	for i := range receiver1.words {
		if receiver1.words[i] != expected1.words[i] {
			t.Fatalf("(*IntSet).AddAll() failed: expected is %v but actual is %v", expected1, receiver1)
		}
	}

	receiver2 := IntSet{[]uint64{0}}
	expected2 := IntSet{[]uint64{0}}

	receiver2.AddAll()

	if len(receiver2.words) != len(expected2.words) {
		t.Fatalf("(*IntSet).AddAll() failed: expected is %v but actual is %v", expected2, receiver2)
	}
	for i := range receiver2.words {
		if receiver2.words[i] != expected2.words[i] {
			t.Fatalf("(*IntSet).AddAll() failed: expected is %v but actual is %v", expected2, receiver2)
		}
	}

	receiver3 := IntSet{[]uint64{0}}
	expected3 := IntSet{[]uint64{0, 1}}

	receiver3.AddAll(64)

	if len(receiver3.words) != len(expected3.words) {
		t.Fatalf("(*IntSet).AddAll() failed: expected is %v but actual is %v", expected3, receiver3)
	}
	for i := range receiver3.words {
		if receiver3.words[i] != expected3.words[i] {
			t.Fatalf("(*IntSet).AddAll() failed: expected is %v but actual is %v", expected3, receiver3)
		}
	}
}

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
			t.Fatalf("(*IntSet).Len() failed: input is %d but actual is %d", test.expected, actual)
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
