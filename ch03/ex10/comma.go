// Copyright 2017 Ken Mirua
package ex10

import (
	"bytes"
)

// comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i > 2 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[len(s)-i-1])
	}

	bytesInBuf := buf.Bytes()
	ret := make([]byte, len(bytesInBuf))
	i := 0
	for j := len(bytesInBuf) - 1; j >= 0; j-- {
		ret[i] = bytesInBuf[j]
		i++
	}
	return string(ret)
}
