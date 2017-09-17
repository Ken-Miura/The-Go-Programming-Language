// Copyright 2017 Ken Miura
package ex11_test

import (
	"fmt"

	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch12/ex11"
)

type Data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
}

var tests = []struct {
	v interface{}
}{
	{&Data{}},
	{&Data{Labels: []string{"golang", "programming"}, MaxResults: 0, Exact: false}},
	{&Data{Labels: []string{"golang", "programming"}, MaxResults: 100, Exact: false}},
	{&Data{Labels: []string{"golang", "programming"}, MaxResults: 0, Exact: true}},
}

// mapをイテレートして文字列出力を作っているため、目視で出力確認
func TestPack(t *testing.T) {
	for _, test := range tests {
		gotQuery := ex11.Pack(test.v)
		fmt.Println(gotQuery)
	}
}
