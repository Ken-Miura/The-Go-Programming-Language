// Copyright 2017 Ken Miura
package ex02_test

import (
	"math/rand"
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch06/ex05"
)

const maxNumOfElems = 1000000
const randMax = 4096
const seed = 1023

func init() {
	rand.Seed(seed)
}

func BenchmarkIntSet_UnionWith(t *testing.B) {
	set1 := ex05.NewIntSet()
	set2 := ex05.NewIntSet()
	for i := 0; i < maxNumOfElems; i++ {
		set1.Add(rand.Intn(randMax))
		set2.Add(rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		set1.UnionWith(set2)
	}
}

func BenchmarkSet_UnionWith(t *testing.B) {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for i := 0; i < maxNumOfElems; i++ {
		set1[rand.Intn(randMax)] = true
		set2[rand.Intn(randMax)] = true
	}
	for i := 0; i < t.N; i++ {
		for i := range set2 {
			_, ok := set1[i]
			if !ok {
				set1[i] = true
			}
		}
	}
}

func BenchmarkIntSet_IntersectWith(t *testing.B) {
	set1 := ex05.NewIntSet()
	set2 := ex05.NewIntSet()
	for i := 0; i < maxNumOfElems; i++ {
		set1.Add(rand.Intn(randMax))
		set2.Add(rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		set1.IntersectWith(set2)
	}
}

func BenchmarkSet_IntersectWith(t *testing.B) {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for i := 0; i < maxNumOfElems; i++ {
		set1[rand.Intn(randMax)] = true
		set2[rand.Intn(randMax)] = true
	}
	for i := 0; i < t.N; i++ {
		for i := range set1 {
			_, ok := set2[i]
			if !ok {
				delete(set1, i)
			}
		}
	}
}

func BenchmarkIntSet_DifferenceWith(t *testing.B) {
	set1 := ex05.NewIntSet()
	set2 := ex05.NewIntSet()
	for i := 0; i < maxNumOfElems; i++ {
		set1.Add(rand.Intn(randMax))
		set2.Add(rand.Intn(randMax))
	}

	for i := 0; i < t.N; i++ {
		set1.DifferenceWith(set2)
	}
}

func BenchmarkSet_DifferenceWith(t *testing.B) {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for i := 0; i < maxNumOfElems; i++ {
		set1[rand.Intn(randMax)] = true
		set2[rand.Intn(randMax)] = true
	}
	for i := 0; i < t.N; i++ {
		for i := range set2 {
			_, ok := set1[i]
			if ok {
				delete(set1, i)
			}
		}
	}
}

func BenchmarkIntSet_SymmetricDifference(t *testing.B) {
	set1 := ex05.NewIntSet()
	set2 := ex05.NewIntSet()
	for i := 0; i < maxNumOfElems; i++ {
		set1.Add(rand.Intn(randMax))
		set2.Add(rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		set1.SymmetricDifference(set2)
	}
}

func BenchmarkSet_SymmetricDifference(t *testing.B) {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)
	for i := 0; i < maxNumOfElems; i++ {
		set1[rand.Intn(randMax)] = true
		set2[rand.Intn(randMax)] = true
	}
	for i := 0; i < t.N; i++ {
		for i := range set1 {
			_, ok := set2[i]
			if ok {
				delete(set1, i)
			}
		}
	}
}

func BenchmarkIntSet_AddAll(t *testing.B) {
	set1 := ex05.NewIntSet()
	integers := make([]int, maxNumOfElems)
	for i := 0; i < maxNumOfElems; i++ {
		integers = append(integers, rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		set1.AddAll(integers...)
	}
}

func BenchmarkSet_AddAll(t *testing.B) {
	set1 := make(map[int]bool)
	integers := make([]int, maxNumOfElems)
	for i := 0; i < maxNumOfElems; i++ {
		integers = append(integers, rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		for _, integer := range integers {
			set1[integer] = true
		}
	}
}

func BenchmarkIntSet_Len(t *testing.B) {
	set1 := ex05.NewIntSet()
	for i := 0; i < maxNumOfElems; i++ {
		set1.Add(rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		set1.Len()
	}
}

func BenchmarkSet_Len(t *testing.B) {
	set1 := make(map[int]bool)
	for i := 0; i < maxNumOfElems; i++ {
		set1[rand.Intn(randMax)] = true
	}
	for i := 0; i < t.N; i++ {
		_ = len(set1)
	}
}

func BenchmarkIntSet_Remove(t *testing.B) {
	set1 := ex05.NewIntSet()
	for i := 0; i < t.N; i++ {
		set1.Remove(rand.Intn(randMax))
	}
}

func BenchmarkSet_Remove(t *testing.B) {
	set1 := make(map[int]bool)
	for i := 0; i < t.N; i++ {
		delete(set1, rand.Intn(randMax))
	}
}

func BenchmarkIntSet_Clear(t *testing.B) {
	set1 := ex05.NewIntSet()
	for i := 0; i < maxNumOfElems; i++ {
		set1.Add(rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		set1.Clear()
	}
}

func BenchmarkSet_Clear(t *testing.B) {
	set1 := make(map[int]bool)
	for i := 0; i < maxNumOfElems; i++ {
		set1[rand.Intn(randMax)] = true
	}
	for i := 0; i < t.N; i++ {
		for i := range set1 {
			delete(set1, i)
		}
	}
}

func BenchmarkIntSet_Copy(t *testing.B) {
	set1 := ex05.NewIntSet()
	for i := 0; i < maxNumOfElems; i++ {
		set1.Add(rand.Intn(randMax))
	}
	for i := 0; i < t.N; i++ {
		_ = set1.Copy()
	}
}

func BenchmarkSet_Copy(t *testing.B) {
	set1 := make(map[int]bool)
	for i := 0; i < maxNumOfElems; i++ {
		set1[rand.Intn(randMax)] = true
	}
	set2 := make(map[int]bool)
	for i := 0; i < t.N; i++ {
		for i := range set1 {
			set2[i] = true
		}
	}
}

func TestIntSet_UnionWith(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		input    *ex05.IntSet
		expected map[int]bool
	}{
		{ex05.NewIntSet(), ex05.NewIntSet(), map[int]bool{}},
		{ex05.NewIntSet(), ex05.NewIntSet(64), map[int]bool{64: true}},
		{ex05.NewIntSet(0), ex05.NewIntSet(0), map[int]bool{0: true}},
		{ex05.NewIntSet(0, 63, 128), ex05.NewIntSet(0, 64, 127), map[int]bool{0: true, 63: true, 64: true, 127: true, 128: true}},
	}

	for _, test := range tests {
		test.receiver.UnionWith(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*IntSet).Len() in (*IntSet).UnionWith(*IntSet) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*IntSet).Has() in (*IntSet).UnionWith(*IntSet) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_IntersectWith(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		input    *ex05.IntSet
		expected map[int]bool
	}{
		{ex05.NewIntSet(), ex05.NewIntSet(), map[int]bool{}},
		{ex05.NewIntSet(), ex05.NewIntSet(64), map[int]bool{}},
		{ex05.NewIntSet(0), ex05.NewIntSet(0), map[int]bool{0: true}},
		{ex05.NewIntSet(0, 1), ex05.NewIntSet(0, 2), map[int]bool{0: true}},
	}

	for _, test := range tests {
		test.receiver.IntersectWith(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*IntSet).Len() in (*IntSet).IntersectWith(*IntSet) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for integer := range test.expected {
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
		expected map[int]bool
	}{
		{ex05.NewIntSet(), ex05.NewIntSet(), map[int]bool{}},
		{ex05.NewIntSet(), ex05.NewIntSet(64), map[int]bool{}},
		{ex05.NewIntSet(0), ex05.NewIntSet(1), map[int]bool{0: true}},
		{ex05.NewIntSet(2, 66), ex05.NewIntSet(0, 2), map[int]bool{66: true}},
	}

	for _, test := range tests {
		test.receiver.DifferenceWith(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*IntSet).Len() in (*IntSet).DifferenceWith(*IntSet) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for integer := range test.expected {
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
		expected map[int]bool
	}{
		{ex05.NewIntSet(), ex05.NewIntSet(), map[int]bool{}},
		{ex05.NewIntSet(), ex05.NewIntSet(64), map[int]bool{64: true}},
		{ex05.NewIntSet(0), ex05.NewIntSet(0), map[int]bool{}},
		{ex05.NewIntSet(1, 2), ex05.NewIntSet(1, 4), map[int]bool{2: true, 4: true}},
	}

	for _, test := range tests {
		test.receiver.SymmetricDifference(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*IntSet).Len() in (*IntSet).SymmetricDifference(*IntSet) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for integer := range test.expected {
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
		expected map[int]bool
	}{
		{ex05.NewIntSet(), []int{1, 2, 3}, map[int]bool{1: true, 2: true, 3: true}},
		{ex05.NewIntSet(), []int{}, map[int]bool{}},
		{ex05.NewIntSet(), []int{64}, map[int]bool{64: true}},
		{ex05.NewIntSet(0), []int{64}, map[int]bool{0: true, 64: true}},
	}

	for _, test := range tests {
		test.receiver.AddAll(test.input...)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*ex05.IntSet).Len() in (*ex05.IntSet).AddAll(...int) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*ex05.IntSet).Has(int) in *ex05.IntSet).AddAll(...int) failed: receiver is expected to have %d", integer)
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
		expected map[int]bool
	}{
		{ex05.NewIntSet(0), 0, map[int]bool{}},
		{ex05.NewIntSet(1), 1, map[int]bool{}},
		{ex05.NewIntSet(0, 1), 2, map[int]bool{0: true, 1: true}},
		{ex05.NewIntSet(64), 64, map[int]bool{}},
		{ex05.NewIntSet(0, 64, 192), 128, map[int]bool{0: true, 64: true, 192: true}},
	}

	for _, test := range tests {
		test.receiver.Remove(test.input)
		if test.receiver.Len() != len(test.expected) {
			t.Fatalf("(*ex05.IntSet).Len() in (*ex05.IntSet).Remove(int) failed: expected is %d but actual is %d", len(test.expected), test.receiver.Len())
		}
		for integer := range test.expected {
			if !test.receiver.Has(integer) {
				t.Fatalf("(*ex05.IntSet).Has(int) in *ex05.IntSet).Remove(int) failed: receiver is expected to have %d", integer)
			}
		}
	}
}

func TestIntSet_Clear(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		values   map[int]bool
	}{
		{ex05.NewIntSet(), map[int]bool{}},
		{ex05.NewIntSet(1), map[int]bool{1: true}},
		{ex05.NewIntSet(0, 1), map[int]bool{0: true, 1: true}},
		{ex05.NewIntSet(64), map[int]bool{64: true}},
		{ex05.NewIntSet(0, 64, 128, 192), map[int]bool{0: true, 64: true, 128: true, 192: true}},
	}

	for _, test := range tests {
		test.receiver.Clear()
		if test.receiver.Len() != 0 {
			t.Fatalf("(*ex05.IntSet).Clear() failed: receiver is expected to have length 0 but length %d", test.receiver.Len())
		}
		for value := range test.values {
			if test.receiver.Has(value) {
				t.Fatalf("(*ex05.IntSet).Clear() failed: receiver is expected to have no valeu but has %d", value)
			}
		}
	}
}

func TestIntSet_Copy(t *testing.T) {
	tests := []struct {
		receiver *ex05.IntSet
		expected map[int]bool
	}{
		{ex05.NewIntSet(), map[int]bool{}},
		{ex05.NewIntSet(1), map[int]bool{1: true}},
		{ex05.NewIntSet(0, 1), map[int]bool{0: true, 1: true}},
		{ex05.NewIntSet(64), map[int]bool{64: true}},
		{ex05.NewIntSet(0, 64, 128, 192), map[int]bool{0: true, 64: true, 128: true, 192: true}},
	}

	for _, test := range tests {
		actual := test.receiver.Copy()
		if actual.Len() != len(test.expected) {
			t.Fatalf("(*ex05.IntSet).Len() in (*ex05.IntSet).Copy() failed: expected is %d but actual is %d", len(test.expected), actual.Len())
		}
		for integer := range test.expected {
			if !actual.Has(integer) {
				t.Fatalf("(*ex05.IntSet).Has() in (*ex05.IntSet).Copy() failed: actual is expected to have %d", integer)
			}
		}
	}
}

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
