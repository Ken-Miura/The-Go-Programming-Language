// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
}

func (t *tree) String() string {
	var result string
	if t != nil {
		result += t.left.String()
		result += strconv.Itoa(t.value)
		result += ", "
		result += t.right.String()
	}
	return result
}

var _ fmt.Stringer = (*tree)(nil)

// サンプルコードより
type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	fmt.Println(root.String()) // ソート後にツリーの値を表示
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
