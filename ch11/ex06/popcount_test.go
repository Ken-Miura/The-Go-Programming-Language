// Copyright 2017 Ken Miura
package ex06

import (
	"strconv"
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch02/ex04"
	"github.com/Ken-Miura/The-Go-Programming-Language/ch02/ex05"

	"gopl.io/ch2/popcount"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex05.PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex04.PopCountByShifting(0x1234567890ABCDEF)
	}
}

func TestPopCounByClearing(t *testing.T) { // テキストのサンプルコードの関数Popcountと同じ出力をするか確認
	var popcountTests = []struct {
		n        uint64 // input
		expected int    // expected
	}{
		{0x1234567890ABCDEF, popcount.PopCount(0x1234567890ABCDEF)},
		{0, popcount.PopCount(0)},
		{1234567890987654321, popcount.PopCount(1234567890987654321)},
	}

	for _, tt := range popcountTests {
		actual := ex05.PopCountByClearing(tt.n)
		if actual != tt.expected {
			t.Fatalf("expected is " + strconv.Itoa(tt.expected) + " but actual is " + strconv.Itoa(actual))
		}
	}
}

func TestPopCountByShifting(t *testing.T) { // テキストのサンプルコードの関数Popcountと同じ出力をするか確認
	var popcountTests = []struct {
		n        uint64 // input
		expected int    // expected
	}{
		{0x1234567890ABCDEF, popcount.PopCount(0x1234567890ABCDEF)},
		{0, popcount.PopCount(0)},
		{1234567890987654321, popcount.PopCount(1234567890987654321)},
	}

	for _, tt := range popcountTests {
		actual := ex04.PopCountByShifting(tt.n)
		if actual != tt.expected {
			t.Fatalf("expected is " + strconv.Itoa(tt.expected) + " but actual is " + strconv.Itoa(actual))
		}
	}
}
