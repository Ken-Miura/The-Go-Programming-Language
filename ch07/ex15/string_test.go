// Copyright 2017 Ken Miura
package main

import (
	"math"
	"testing"
)

var tests = []struct {
	input    string
	env      Env
	expected Expr
}{
	{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, nil},
	{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, nil},
	{"5 / 9 * (F - 32)", Env{"F": -40}, nil},
	{"min(x, y, 3)", Env{"x": 9, "y": 10}, nil},
}

func init() {
	for i := range tests {
		tests[i].expected, _ = Parse(tests[i].input)
	}
}

func TestString(t *testing.T) {
	for _, test := range tests {
		result, _ := Parse(test.expected.String())
		for result.String() != test.expected.String() {
			t.Errorf("%s is expected but result is %s", test.expected.String(), result.String())
		}
	}
}
