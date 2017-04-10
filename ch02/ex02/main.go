// Copyright 2017 Ken Mirua
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			t, err := strconv.ParseFloat(input.Text(), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "mf: %v\n", err)
				os.Exit(1)
			}
			displayNumInMeterAndFeet(t)
		}
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "mf: %v\n", err)
				os.Exit(1)
			}
			displayNumInMeterAndFeet(t)
		}
	}
}

func displayNumInMeterAndFeet(num float64) {
	f := Feet(num)
	m := Meter(num)
	fmt.Printf("%s = %s, %s = %s\n",
		f, FToM(f), m, MToF(m))
}
