// Copyright 2017 Ken Mirua
package ex03

import (
	"strconv"
	"testing"
)

func TestPopCountUsingLoop(t *testing.T) { // テキストのサンプルコードの関数Popcountと同じ出力をするか確認
	expected1 := PopCount(0x1234567890ABCDEF)
	actual1 := PopCountUsingLoop(0x1234567890ABCDEF)
	if actual1 != expected1 {
		t.Fatalf("expected is " + strconv.Itoa(expected1) + " but actual is " + strconv.Itoa(actual1))
	}

	expected2 := PopCount(0)
	actual2 := PopCountUsingLoop(0)
	if actual2 != expected2 {
		t.Fatalf("expected is " + strconv.Itoa(expected2) + " but actual is " + strconv.Itoa(actual2))
	}

	expected3 := PopCount(1234567890987654321)
	actual3 := PopCountUsingLoop(1234567890987654321)
	if actual3 != expected3 {
		t.Fatalf("expected is " + strconv.Itoa(expected3) + " but actual is " + strconv.Itoa(actual3))
	}
}
