// Copyright 2017 Ken Mirua
package ex06

import (
	"unicode"
)

func CompressSpaces(s []byte) []byte {
	runes := []rune(string(s))
	for i := 0; i < len(runes)-1; {
		if unicode.IsSpace(runes[i]) && unicode.IsSpace(runes[i+1]) {
			runes = remove(runes, i+1)
			continue
		}
		i++
	}
	runes = replaceSpacesToAsciiSpase(runes)
	return []byte(string(runes))
}

func replaceSpacesToAsciiSpase(runes []rune) []rune {
	for i := range runes {
		if unicode.IsSpace(runes[i]) && runes[i] != ' ' {
			runes[i] = ' '
		}
	}
	return runes
}

func remove(slice []rune, i int) []rune {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
