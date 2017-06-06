// Copyright 2017 Ken Miura
package ex04

// TODO もとのスライスを返さないといけない＋ワンパスで処理をしないといけない
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
