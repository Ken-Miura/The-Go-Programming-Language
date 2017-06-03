// Copyright 2017 Ken Mirua
package main

import "testing"

var tests = []struct {
	src     string
	replace func(string) string
	dst     string
}{
	{"This is $foo function test", func(s string) string {
		if s != "foo" {
			return "bag arg"
		}
		return "Expand"
	}, "This is Expand function test"},
}

func TestExpand(t *testing.T) {
	for _, test := range tests {
		result := Expand(test.src, test.replace)
		if result != test.dst {
			t.Fatalf("'%s' is expected but result is '%s'", test.dst, result)
		}
	}
}
