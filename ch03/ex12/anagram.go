// Copyright 2017 Ken Miura
package ex12

import (
	"unicode"
	"unicode/utf8"
)

func IsAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if utf8.RuneCountInString(s1) != utf8.RuneCountInString(s2) {
		return false
	}

	runeCount1 := make(map[rune]int)
	for _, r1 := range s1 {
		r1 = unicode.ToLower(unicode.ToUpper(r1))
		runeCount1[r1]++
	}

	runeCount2 := make(map[rune]int)
	for _, r2 := range s2 {
		r2 = unicode.ToLower(unicode.ToUpper(r2))
		runeCount2[r2]++
	}

	for k1, v1 := range runeCount1 {
		v2, ok := runeCount2[k1]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}
