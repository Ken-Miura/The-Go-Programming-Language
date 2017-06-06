// Copyright 2017 Ken Miura
package ex11

import (
	"bytes"
	"strings"
)

// comma inserts commas in a floating point number string.
func Comma(s string) string {
	period := strings.LastIndex(s, ".")
	if period == -1 {
		return commaToSignedInteger(s)
	} else {
		return commaToSignedInteger(s[:period]) + s[period:]
	}
}

func commaToSignedInteger(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i != 0 && i%3 == 0 && (s[len(s)-i-1] != '+' || s[len(s)-i-1] != '-') {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[len(s)-i-1])
	}

	bytesInBuf := buf.Bytes()
	for i, j := 0, len(bytesInBuf)-1; i < j; i, j = i+1, j-1 {
		bytesInBuf[i], bytesInBuf[j] = bytesInBuf[j], bytesInBuf[i]
	}
	return string(bytesInBuf)
}
