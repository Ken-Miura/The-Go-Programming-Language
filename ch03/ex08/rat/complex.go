// Copyright 2017 Ken Mirua
package main

import (
	"math/big"
)

type Complex struct {
	re *big.Rat
	im *big.Rat
}

func NewComplex(re, im *big.Rat) *Complex {
	return &Complex{re, im}
}

func Add(a, b *Complex) *Complex {
	tempRe := big.NewRat(0, 1)
	tempRe.Add(a.re, b.re)

	tempImag := big.NewRat(0, 1)
	tempImag.Add(a.im, b.im)

	return &Complex{tempRe, tempImag}
}

func Multiply(a, b *Complex) *Complex {
	temp1 := big.NewRat(0, 1)
	temp1.Mul(a.re, b.re)

	temp2 := big.NewRat(0, 1)
	temp2.Mul(a.im, b.im)

	tempRe := big.NewRat(0, 1)
	tempRe.Sub(temp1, temp2)

	temp3 := big.NewRat(0, 1)
	temp3.Mul(a.re, b.im)

	temp4 := big.NewRat(0, 1)
	temp4.Mul(a.im, b.re)

	tempImag := big.NewRat(0, 1)
	tempImag.Add(temp3, temp4)

	return &Complex{tempRe, tempImag}
}

func SquaredAbs(bc *Complex) *big.Rat {
	tempRe := big.NewRat(0, 1)
	tempRe.Mul(bc.re, bc.re)

	tempImag := big.NewRat(0, 1)
	tempImag.Mul(bc.im, bc.im)

	ret := big.NewRat(0, 1)
	return ret.Add(tempRe, tempImag)
}

// TODO 全部解いた後時間があれば
//func Abs(bc *Complex) *big.Rat {
//	temp := squaredAbs(bc)
//	return sqrt(temp)
//}
//
//var tolerance *big.Rat = big.NewRat(1, 100000000000000)
//
//func sqrt(s *big.Rat) *big.Rat {
//	// Babylonian methodのつもり
//
//	x := big.NewRat(0, 1)
//	x.Set(s)
//
//	lastX := big.NewRat(0, 1)
//
//	two := big.NewRat(2, 1)
//
//	for big.NewRat(0, 1).Abs(big.NewRat(0, 1).Sub(x, lastX)).Cmp(tolerance) > 0 {
//
//		fmt.Println(lastX, x)
//
//		lastX.Set(x)
//
//		temp1 := big.NewRat(0, 1)
//		temp1.Quo(s, x)
//
//		temp2 := big.NewRat(0, 1)
//		temp2.Add(x, temp1)
//
//		temp3 := big.NewRat(0, 1)
//		temp3.Quo(temp2, two)
//
//		x.Set(temp3)
//	}
//	return x
//}
