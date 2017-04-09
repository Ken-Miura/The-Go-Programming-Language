// Copyright 2017 Ken Mirua
package ex01

import "testing"

func TestCToF(t *testing.T) {
	expected1 := AbsoluteZeroF
	actual1 := CToF(AbsoluteZeroC)
	if actual1 != expected1 {
		t.Fatal("Failed conversion F to C")
	}

	expected2 := FreezingF
	actual2 := CToF(FreezingC)
	if actual2 != expected2 {
		t.Fatal("Failed conversion F to C")
	}

	expected3 := BoilingF
	actual3 := CToF(BoilingC)
	if actual3 != expected3 {
		t.Fatal("Failed conversion F to C")
	}
}

func TestFToC(t *testing.T) {
	expected1 := AbsoluteZeroC
	actual1 := FToC(AbsoluteZeroF)
	if actual1 != expected1 {
		t.Fatal("Failed conversion C to F")
	}

	expected2 := FreezingC
	actual2 := FToC(FreezingF)
	if actual2 != expected2 {
		t.Fatal("Failed conversion C to F")
	}

	expected3 := BoilingC
	actual3 := FToC(BoilingF)
	if actual3 != expected3 {
		t.Fatal("Failed conversion C to F")
	}
}
