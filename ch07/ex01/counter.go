// Copyright 2017 Ken Miura
package ex01

import (
	"bufio"
	"bytes"
	"io"
)

type WordCounter struct {
	wordCount int
}

func (c *WordCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	input.Split(bufio.ScanWords)
	tmpWordCount := 0
	for input.Scan() {
		if err := input.Err(); err != nil {
			return 0, err
		}
		tmpWordCount++
	}
	c.wordCount = c.wordCount + tmpWordCount
	return len(p), nil
}

func (c *WordCounter) WordCount() int {
	return c.wordCount
}

var _ io.Writer = (*WordCounter)(nil)

type LineCounter struct {
	lineCount int
}

func (c *LineCounter) Write(p []byte) (int, error) {
	input := bufio.NewScanner(bytes.NewReader(p))
	tmpLineCount := 0
	for input.Scan() {
		if err := input.Err(); err != nil {
			return 0, err
		}
		tmpLineCount++
	}
	c.lineCount = c.lineCount + tmpLineCount
	return len(p), nil
}

func (c *LineCounter) LineCount() int {
	return c.lineCount
}

var _ io.Writer = (*LineCounter)(nil)
