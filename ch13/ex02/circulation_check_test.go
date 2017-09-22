// Copyright 2017 Ken Miura
package ex02_test

import (
	"testing"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch13/ex02"
)

func TestIsCircularDataStructure(t *testing.T) {
	type CyclePtr *CyclePtr
	var cyclePtr1 CyclePtr
	cyclePtr1 = &cyclePtr1

	if !ex02.IsCyclic(cyclePtr1) {
		t.Errorf("IsCyclic(cyclePtr1) returned false")
	}

	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	if !ex02.IsCyclic(a) {
		t.Errorf("IsCyclic(a) returned false")
	}

	if !ex02.IsCyclic(b) {
		t.Errorf("IsCyclic(b) returned false")
	}

	if !ex02.IsCyclic(c) {
		t.Errorf("IsCyclic(c) returned false")
	}

	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	if ex02.IsCyclic(strangelove) {
		t.Errorf("IsCyclic(strangelove) returned true")
	}
}
