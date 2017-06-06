// Copyright 2017 Ken Miura
package main

import (
	"math/big"
	"testing"
)

var a1 = NewComplex(big.NewRat(1, 1), big.NewRat(2, 1))
var b1 = NewComplex(big.NewRat(3, 1), big.NewRat(4, 1))
var addexpected1 = NewComplex(big.NewRat(4, 1), big.NewRat(6, 1))
var multiplyexpected1 = NewComplex(big.NewRat(-5, 1), big.NewRat(10, 1))
var squaredabsexpected1 = big.NewRat(5, 1)

var a2 = NewComplex(big.NewRat(1, 1), big.NewRat(0, 1))
var b2 = NewComplex(big.NewRat(3, 1), big.NewRat(0, 1))
var addexpected2 = NewComplex(big.NewRat(4, 1), big.NewRat(0, 1))
var multiplyexpected2 = NewComplex(big.NewRat(3, 1), big.NewRat(0, 1))
var squaredabsexpected2 = big.NewRat(1, 1)

var a3 = NewComplex(big.NewRat(0, 1), big.NewRat(2, 1))
var b3 = NewComplex(big.NewRat(0, 1), big.NewRat(4, 1))
var addexpected3 = NewComplex(big.NewRat(0, 1), big.NewRat(6, 1))
var multiplyexpected3 = NewComplex(big.NewRat(-8, 1), big.NewRat(0, 1))
var squaredabsexpected3 = big.NewRat(4, 1)

var a4 = NewComplex(big.NewRat(0, 1), big.NewRat(0, 1))
var b4 = NewComplex(big.NewRat(0, 1), big.NewRat(0, 1))
var addexpected4 = NewComplex(big.NewRat(0, 1), big.NewRat(0, 1))
var multiplyexpected4 = NewComplex(big.NewRat(0, 1), big.NewRat(0, 1))
var squaredabsexpected4 = big.NewRat(0, 1)

var tests = []struct {
	a                  *Complex // input
	b                  *Complex // input
	addexpected        *Complex // add func expected
	multiplyexpected   *Complex // multiply func expected
	squaredabsexpected *big.Rat // abs func expected
}{
	{a1, b1, addexpected1, multiplyexpected1, squaredabsexpected1},
	{a2, b2, addexpected2, multiplyexpected2, squaredabsexpected2},
	{a3, b3, addexpected3, multiplyexpected3, squaredabsexpected3},
	{a4, b4, addexpected4, multiplyexpected4, squaredabsexpected4},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		actual := Add(tt.a, tt.b)
		result1 := actual.re.Cmp(tt.addexpected.re)
		result2 := actual.im.Cmp(tt.addexpected.im)
		if result1 != 0 || result2 != 0 {
			t.Fatalf("expected is %v + (%v)i, but actual is %v + (%v)i", tt.addexpected.re, tt.addexpected.im, actual.re, actual.im)
		}

	}
}

func TestMultiply(t *testing.T) {
	for _, tt := range tests {
		actual := Multiply(tt.a, tt.b)
		result1 := actual.re.Cmp(tt.multiplyexpected.re)
		result2 := actual.im.Cmp(tt.multiplyexpected.im)
		if result1 != 0 || result2 != 0 {
			t.Fatalf("expected is %v + (%v)i, but actual is %v + (%v)i", tt.multiplyexpected.re, tt.multiplyexpected.im, actual.re, actual.im)
		}

	}
}

func TestSquredAbs(t *testing.T) {
	for _, tt := range tests {
		actual := SquaredAbs(tt.a)
		if actual.Cmp(tt.squaredabsexpected) != 0 {
			t.Fatalf("expected is %v, but actual is %v", tt.squaredabsexpected, actual)
		}
	}
}
