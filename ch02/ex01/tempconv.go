// Copyright 2017 Ken Mirua
// Package ex01 performs Celsius and Fahrenheit conversions.
package ex01

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

const (
	AbsoluteZeroF Fahrenheit = -459.66999999999996
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
