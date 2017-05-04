// Copyright 2017 Ken Mirua
package main

import (
	"math/big"
)

type BigComplex struct {
	re *big.Float
	im *big.Float
}

func NewBigComplex(re, im *big.Float) (*BigComplex, bool) {
	if re.Prec() != im.Prec() {
		return &BigComplex{}, false
	}
	if re.Mode() != im.Mode() {
		return &BigComplex{}, false
	}
	return &BigComplex{re, im}, true
}

func Add(a, b *BigComplex) *BigComplex {
	tempRe := big.NewFloat(0.0)
	tempRe.SetPrec(a.re.Prec())
	tempRe.SetMode(a.re.Mode())
	tempRe.Add(a.re, b.re)

	tempImag := big.NewFloat(0.0)
	tempImag.SetPrec(a.im.Prec())
	tempImag.SetMode(a.im.Mode())
	tempImag.Add(a.im, b.im)

	return &BigComplex{tempRe, tempImag}
}

func Multiply(a, b *BigComplex) *BigComplex {
	tempRe := big.NewFloat(0.0)
	tempRe.SetPrec(a.re.Prec())
	tempRe.SetMode(a.re.Mode())
	tempRe.Mul(a.re, b.re)

	tempImag := big.NewFloat(0.0)
	tempImag.SetPrec(a.im.Prec())
	tempImag.SetMode(a.im.Mode())
	tempImag.Mul(a.im, b.im)

	return &BigComplex{tempRe, tempImag}
}

func Abs(bc *BigComplex) *big.Float {
	temp := squaredAbs(bc)
	return sqrt(temp)
}

func squaredAbs(bc *BigComplex) *big.Float {
	tempRe := big.NewFloat(0.0)
	tempRe.SetPrec(bc.re.Prec())
	tempRe.SetMode(bc.re.Mode())
	tempRe.Mul(bc.re, bc.re)

	tempImag := big.NewFloat(0.0)
	tempImag.SetPrec(bc.im.Prec())
	tempImag.SetMode(bc.im.Mode())
	tempImag.Mul(bc.im, bc.im)

	ret := big.NewFloat(0.0)
	ret.SetPrec(bc.re.Prec())
	ret.SetMode(bc.re.Mode())

	return ret.Add(tempRe, tempImag)
}

func sqrt(s *big.Float) *big.Float {
	// Babylonian methodのつもり

	x := big.NewFloat(0.0)
	x.SetPrec(s.Prec())
	x.SetMode(s.Mode())
	x.Copy(s)

	lastX := big.NewFloat(0.0)
	lastX.SetPrec(s.Prec())
	lastX.SetMode(s.Mode())

	two := big.NewFloat(2.0)
	two.SetPrec(s.Prec())
	two.SetMode(s.Mode())

	for x.Cmp(lastX) != 0 {
		lastX.Copy(x)

		temp1 := big.NewFloat(0.0)
		temp1.SetPrec(s.Prec())
		temp1.SetMode(s.Mode())
		temp1.Quo(s, x)

		temp2 := big.NewFloat(0.0)
		temp2.SetPrec(s.Prec())
		temp2.SetMode(s.Mode())
		temp2.Add(x, temp1)

		temp3 := big.NewFloat(0.0)
		temp3.SetPrec(s.Prec())
		temp3.SetMode(s.Mode())
		temp3.Quo(temp2, two)

		x.Copy(temp3)
	}
	return x
}
