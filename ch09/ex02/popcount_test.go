// Copyright 2017 Ken Miura
package ex02_test

import (
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch09/ex02"
	"gopl.io/ch2/popcount"
)

func TestPopCountUsingLoop(t *testing.T) { // テキストのサンプルコードの関数Popcountと同じ出力をするか確認
	var popcountTests = []struct {
		n        uint64 // input
		expected int    // expected
	}{
		{0x1234567890ABCDEF, popcount.PopCount(0x1234567890ABCDEF)},
		{0, popcount.PopCount(0)},
		{1234567890987654321, popcount.PopCount(1234567890987654321)},
	}

	for _, tt := range popcountTests {
		actual := ex02.PopCount(tt.n)
		if actual != tt.expected {
			t.Fatalf("expected is %d but actual is %d", tt.expected, actual)
		}
	}
}
