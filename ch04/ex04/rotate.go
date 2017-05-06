// Copyright 2017 Ken Mirua
package ex04

func Rotate(s []int, n int) []int {
	result := make([]int, len(s))
	if n < 0 {
		for i := range s {
			j := i + n
			if j < 0 {
				j = len(s) + j
			}
			result[j] = s[i]
		}
	} else {
		for i := range s {
			j := (i + n) % len(s)
			result[j] = s[i]
		}
	}
	return result
}
