// Copyright 2017 Ken Mirua
package main

import "fmt"

type Meter float64
type Feet float64

const (
	OneFeetInMeters Meter = 0.30480370641307
)

func (m Meter) String() string { return fmt.Sprintf("%g m", m) }
func (f Feet) String() string  { return fmt.Sprintf("%g ft", f) }

func MToF(m Meter) Feet { return Feet(3.2808 * m) }

func FToM(f Feet) Meter { return Meter(f / 3.2808) }
