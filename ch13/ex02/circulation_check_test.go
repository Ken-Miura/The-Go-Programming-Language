// Copyright 2017 Ken Miura
package ex02_test

import (
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch13/ex02"
)

func TestIsCyclic(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b, c, d, e := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}, &link{value: "d"}, &link{value: "e"}
	a.tail, b.tail = b, a
	c.tail = c
	d.tail = e

	tests := []struct {
		input *link
		want  bool
	}{
		{a, true},
		{b, true},
		{c, true},
		{d, false},
	}

	for _, test := range tests {
		got := ex02.IsCyclic(test.input)
		if got != test.want {
			t.Errorf("IsCyclic returned %t", got)
		}
	}
}
