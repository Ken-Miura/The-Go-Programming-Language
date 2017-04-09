// Copyright 2017 Ken Mirua
package main

import "testing"

func TestMToF(t *testing.T) {

	var expected1 Feet = 1
	actual1 := MToF(OneFeetInMeters)
	if actual1 != expected1 {
		t.Fatal("Failed conversion meter to feet")
	}

	var expected2 Feet = 0
	actual2 := MToF(0)
	if actual2 != expected2 {
		t.Fatal("Failed conversion meter to feet")
	}

	var expected3 Feet = 3.2808
	actual3 := MToF(1)
	if actual3 != expected3 {
		t.Fatal("Failed conversion meter to feet")
	}
}

func TestFToM(t *testing.T) {
	expected1 := OneFeetInMeters
	actual1 := FToM(1)
	if actual1 != expected1 {
		t.Fatal("Failed conversion feet to meter")
	}

	var expected2 Meter = 0
	actual2 := FToM(0)
	if actual2 != expected2 {
		t.Fatal("Failed conversion feet to meter")
	}

	var expected3 Meter = 1
	actual3 := FToM(3.2808)
	if actual3 != expected3 {
		t.Fatal("Failed conversion feet to meter")
	}
}
