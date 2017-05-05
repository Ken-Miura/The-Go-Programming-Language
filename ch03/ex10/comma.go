// Copyright 2017 Ken Mirua
package ex10

import (
	"bytes"
)

// comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i != 0 && i%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[len(s)-i-1])
	}

	bytesInBuf := buf.Bytes()
	i, j := 0, len(bytesInBuf)-1
	for i < j {
		bytesInBuf[i], bytesInBuf[j] = bytesInBuf[j], bytesInBuf[i]
		i++
		j--
	}
	return string(bytesInBuf)
}
