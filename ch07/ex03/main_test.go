// Copyright 2017 Ken Mirua
package main

import (
	"math/rand"
	"sort"
	"testing"
)

// サンプルコードからほぼそのまま引用
// Sort関数内に修正（Stringメソッドを利用してツリー内の値表示）を入れたので、一応Sortのテストが通ることを確認しておく。
func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}
