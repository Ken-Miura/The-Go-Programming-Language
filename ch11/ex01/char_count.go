// Copyright 2017 Ken Miura
// gopl.io/ch4/charcountより単体テストできるように必要な部分を抽出
package ex01

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"unicode/utf8"
)

// CharCount returns counts of Unicode characters, counts of lengths of UTF-8 encodings and counts of invalid UTF-8 characters in r.
func CharCount(r io.Reader) (map[rune]int, [utf8.UTFMax + 1]int, int, error) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, [utf8.UTFMax + 1]int{}, 0, fmt.Errorf("charcount: %v", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return counts, utflen, invalid, nil
}
