// Copyright 2017 Ken Mirua
package main

import "fmt"

func main() {

	func() {
		defer func() {
			p := recover()
			if p != nil {
				fmt.Println(p)
			}
		}()
		min()
	}()

	func() {
		defer func() {
			p := recover()
			if p != nil {
				fmt.Println(p)
			}
		}()
		max()
	}()
}

func min(values ...int) int {
	if len(values) == 0 {
		panic("min: cannot calculate due to no argument")
	}
	return Min(values[0], values[1:]...)
}

func max(values ...int) int {
	if len(values) == 0 {
		panic("max: cannot calculate due to no argument")
	}
	return Max(values[0], values[1:]...)
}

func Min(val int, values ...int) int {
	result := val
	for _, v := range values {
		if result > v {
			result = v
		}
	}
	return result
}

func Max(val int, values ...int) int {
	result := val
	for _, v := range values {
		if result < v {
			result = v
		}
	}
	return result
}
