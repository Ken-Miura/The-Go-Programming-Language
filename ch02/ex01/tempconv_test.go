// Copyright 2017 Ken Mirua
package ex01

import (
	"math"
	"testing"
)

var testsBetweenKelvinAndFahrenheit = []struct {
	kelvin     Kelvin
	fahrenheit Fahrenheit
}{
	{0, -459.67},
	{273.15, 32},
	{373.15, 212},
}

var tolerance float64 = 0.001

var testsBetweenKelvinAndCelsius = []struct {
	kelvin  Kelvin
	celsius Celsius
}{
	{0, -273.15},
	{273.15, 0},
	{373.15, 100},
}

func TestKToF(t *testing.T) {
	for _, tt := range testsBetweenKelvinAndFahrenheit {
		actual := KToF(tt.kelvin)
		diff := math.Abs(float64(tt.fahrenheit - actual))
		if diff > tolerance {
			t.Fatal("Failed conversion K to F")
		}
	}
}

func TestFToK(t *testing.T) {
	for _, tt := range testsBetweenKelvinAndFahrenheit {
		actual := FToK(tt.fahrenheit)
		diff := math.Abs(float64(tt.kelvin - actual))
		if diff > tolerance {
			t.Fatal("Failed conversion F to K")
		}
	}
}

func TestKToC(t *testing.T) {
	for _, tt := range testsBetweenKelvinAndCelsius {
		actual := KToC(tt.kelvin)
		if actual != tt.celsius {
			t.Fatal("Failed conversion K to C")
		}
	}
}

func TestCToK(t *testing.T) {
	for _, tt := range testsBetweenKelvinAndCelsius {
		actual := CToK(tt.celsius)
		if actual != tt.kelvin {
			t.Fatal("Failed conversion C to K")
		}
	}
}
