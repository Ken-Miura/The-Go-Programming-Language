// Copyright 2017 Ken Miura
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
		dollarIndex := strings.Index(s, "$")
		if dollarIndex < 0 {
			result += s
			return result
		}
		spaceIndex := len(s)
		for i := dollarIndex; i < len(s); {
			r, num := utf8.DecodeRune([]byte(s[i:]))
			if unicode.IsSpace(r) {
				spaceIndex = i
				break
			}
			i += num
		}
		replacement := f(s[dollarIndex+1 : spaceIndex])
		result += s[:dollarIndex]
		result += replacement
		s = s[spaceIndex:]
	}
	return result
}

func main() {

}
