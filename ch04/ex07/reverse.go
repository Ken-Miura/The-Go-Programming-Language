// Copyright 2017 Ken Miura
package ex07

import "unicode/utf8"

func Reverse(s []byte) []byte {
	size := 0
	for size < len(s) {
		_, runeLength := utf8.DecodeRune(s[size:])
		for i, j := size, (size+runeLength)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		size += runeLength
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
