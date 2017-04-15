// Copyright 2017 Ken Mirua
package ex04

import (
	"strconv"
	"testing"
)

func TestPopCountUsingLoop(t *testing.T) { // テキストのサンプルコードの関数Popcountと同じ出力をするか確認
	var popcountTests = []struct {
		n        uint64 // input
		expected int    // expected
	}{
		{0x1234567890ABCDEF, PopCount(0x1234567890ABCDEF)},
		{0, PopCount(0)},
		{1234567890987654321, PopCount(1234567890987654321)},
	}

	for _, tt := range popcountTests {
		actual := PopCountByShifting(tt.n)
		if actual != tt.expected {
			t.Fatalf("expected is " + strconv.Itoa(tt.expected) + " but actual is " + strconv.Itoa(actual))
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}
