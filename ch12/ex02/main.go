// Copyright 2017 Ken Miura
package main

func main() {
	// 連取問題12の2
	// a pointer that points to itself
	type P *P
	var p P
	p = &p
	Display("p", p)

	// a map that contains itself
	type M map[string]M
	m := make(M)
	m[""] = m
	Display("m", m)

	// a slice that contains itself
	type S []S
	s := make(S, 1)
	s[0] = s
	Display("s", s)

	// a linked list that eats its own tail
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}
	Display("c", c)

	// 練習問題12の1
	// DisplayUsingStructAsKey()
	// DisplayUsingArrayAsKey()
}

func DisplayUsingStructAsKey() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
	}
	movieSet := make(map[Movie]bool)
	movieSet[strangelove] = true
	Display("map", movieSet)
}

const arrayLength = 5

func DisplayUsingArrayAsKey() {
	var intArray [arrayLength]int
	for i := 0; i < arrayLength; i++ {
		intArray[i] = i
	}
	intArraySet := make(map[[arrayLength]int]bool)
	intArraySet[intArray] = true
	Display("intArraySet", intArraySet)
}
