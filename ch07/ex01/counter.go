// Copyright 2017 Ken Miura
package ex01

import (
	"bufio"
	"bytes"
	"io"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	count := 0
	for input.Scan() {
		count++
	}
	return count, nil
}

var _ io.Writer = (*WordCounter)(nil)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	count := 0
	for input.Scan() {
		count++
	}
	return count, nil
}

var _ io.Writer = (*LineCounter)(nil)
