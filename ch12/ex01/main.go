// Copyright 2017 Ken Miura
package main

func main() {
	DisplayUsingStructAsKey()
	DisplayUsingArrayAsKey()
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
