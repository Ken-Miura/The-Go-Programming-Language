// Copyright 2017 Ken Miura
// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type letterClassification int

const (
	letter letterClassification = iota
	notLetter
)

type numberClassification int

const (
	number numberClassification = iota
	notNumber
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	var letterClassification [2]int // count of Unicode letter and not Unicode letter characters
	var numberClassification [2]int // count of Unicode number and not Unicode number characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		if unicode.IsLetter(r) {
			letterClassification[letter]++
		} else {
			letterClassification[notLetter]++
		}
		if unicode.IsNumber(r) {
			numberClassification[number]++
		} else {
			numberClassification[notNumber]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

	fmt.Print("\nletter\t\tcount\n")
	fmt.Printf("letter\t\t%d\n", letterClassification[letter])
	fmt.Printf("not letter\t%d\n", letterClassification[notLetter])

	fmt.Print("\nnumber\t\tcount\n")
	fmt.Printf("number\t\t%d\n", numberClassification[number])
	fmt.Printf("not number\t%d\n", numberClassification[notNumber])
}
