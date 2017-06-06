// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"os"
)

func main() {
	s := withoutReturn()
	if s == "" {
		fmt.Printf("withoutReturn returns zero value: %s\n", s)
		os.Exit(1)
	}
	fmt.Printf("withoutReturn returns non zero value: %s\n", s)
}

func withoutReturn() (result string) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case bailout{}:
			result = "non zero value of type string"
		default:
			panic(p)
		}
	}()
	panic(bailout{})
}
