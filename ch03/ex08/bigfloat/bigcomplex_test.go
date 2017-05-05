// Copyright 2017 Ken Mirua
package main

import (
	"math/big"
	"testing"
)

func TestNewBigComplex(t *testing.T) {
	x1 := big.NewFloat(1.0)
	y1 := big.NewFloat(2.0)
	z1, ok1 := NewBigComplex(x1, y1)
	if !ok1 {
		t.Fatalf("expected is true, but actual is %t", ok1)
	}
	if z1.re.Cmp(x1) != 0 || z1.im.Cmp(y1) != 0 {
		t.Fatalf("expected is %g + (%g)i, but actual is %g + (%g)i", x1, y1, z1.re, z1.im)
	}

	x2 := big.NewFloat(1.0)
	x2.SetPrec(70)
	y2 := big.NewFloat(2.0)
	y2.SetPrec(90)
	_, ok2 := NewBigComplex(x2, y2)
	if ok2 {
		t.Fatalf("expected is false, but actual is %t", ok2)
	}

	x3 := big.NewFloat(1.0)
	x3.SetMode(big.ToNearestEven)
	y3 := big.NewFloat(2.0)
	y3.SetMode(big.ToNearestAway)
	_, ok3 := NewBigComplex(x3, y3)
	if ok3 {
		t.Fatalf("expected is false, but actual is %t", ok3)
	}
}

var a1, _ = NewBigComplex(big.NewFloat(1.0), big.NewFloat(2.0))
var b1, _ = NewBigComplex(big.NewFloat(3.0), big.NewFloat(4.0))
var addexpected1, _ = NewBigComplex(big.NewFloat(4.0), big.NewFloat(6.0))
var multiplyexpected1, _ = NewBigComplex(big.NewFloat(-5.0), big.NewFloat(10.0))
var absexpected1 = big.NewFloat(2.23606797749979)

var a2, _ = NewBigComplex(big.NewFloat(1.0), big.NewFloat(0.0))
var b2, _ = NewBigComplex(big.NewFloat(3.0), big.NewFloat(0.0))
var addexpected2, _ = NewBigComplex(big.NewFloat(4.0), big.NewFloat(0.0))
var multiplyexpected2, _ = NewBigComplex(big.NewFloat(3.0), big.NewFloat(0.0))
var absexpected2 = big.NewFloat(1.0)

var a3, _ = NewBigComplex(big.NewFloat(0.0), big.NewFloat(2.0))
var b3, _ = NewBigComplex(big.NewFloat(0.0), big.NewFloat(4.0))
var addexpected3, _ = NewBigComplex(big.NewFloat(0.0), big.NewFloat(6.0))
var multiplyexpected3, _ = NewBigComplex(big.NewFloat(-8.0), big.NewFloat(0.0))
var absexpected3 = big.NewFloat(2.0)

var a4, _ = NewBigComplex(big.NewFloat(0.0), big.NewFloat(0.0))
var b4, _ = NewBigComplex(big.NewFloat(0.0), big.NewFloat(0.0))
var addexpected4, _ = NewBigComplex(big.NewFloat(0.0), big.NewFloat(0.0))
var multiplyexpected4, _ = NewBigComplex(big.NewFloat(0.0), big.NewFloat(0.0))
var absexpected4 = big.NewFloat(0.0)

var tests = []struct {
	a                *BigComplex // input
	b                *BigComplex // input
	addexpected      *BigComplex // add func expected
	multiplyexpected *BigComplex // multiply func expected
	absexpected      *big.Float  // abs func expected
}{
	{a1, b1, addexpected1, multiplyexpected1, absexpected1},
	{a2, b2, addexpected2, multiplyexpected2, absexpected2},
	{a3, b3, addexpected3, multiplyexpected3, absexpected3},
	{a4, b4, addexpected4, multiplyexpected4, absexpected4},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		actual := Add(tt.a, tt.b)
		result1 := actual.re.Cmp(tt.addexpected.re)
		result2 := actual.im.Cmp(tt.addexpected.im)
		if result1 != 0 || result2 != 0 {
			t.Fatalf("expected is %g + (%g)i, but actual is %g + (%g)i", tt.addexpected.re, tt.addexpected.im, actual.re, actual.im)
		}

	}
}

func TestMultiply(t *testing.T) {
	for _, tt := range tests {
		actual := Multiply(tt.a, tt.b)
		result1 := actual.re.Cmp(tt.multiplyexpected.re)
		result2 := actual.im.Cmp(tt.multiplyexpected.im)
		if result1 != 0 || result2 != 0 {
			t.Fatalf("expected is %g + (%g)i, but actual is %g + (%g)i", tt.multiplyexpected.re, tt.multiplyexpected.im, actual.re, actual.im)
		}

	}
}

func TestAbs(t *testing.T) {
	for _, tt := range tests {
		actual := Abs(tt.a)
		if actual.Cmp(tt.absexpected) != 0 {
			t.Fatalf("expected is %g, but actual is %g", tt.absexpected, actual)
		}

	}
}
