// Copyright 2017 Ken Mirua
package main

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// TODO regexpのExpandメソッドで簡単にできそう？
func Expand(s string, f func(string) string) string {
	result := ""
	for len(s) > 0 {
		index := strings.Index(s, "$")
		if index < 0 {
			result += s
			break
		}
		spaceIndex := len(s)
		for i := index; i < len(s); {
			r, num := utf8.DecodeRune([]byte(s[i:]))
			if unicode.IsSpace(r) {
				spaceIndex = i
				break
			}
			i += num
		}
		replacement := f(s[index+1 : spaceIndex])
		result += s[:index]
		result += replacement
		s = s[spaceIndex:]
	}
	return result
}

func main() {

}
