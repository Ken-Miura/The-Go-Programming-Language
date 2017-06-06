// Copyright 2017 Ken Miura
package ex02

import "io"

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	counter := byteCounter{w, 0}
	return &counter, &(counter.count)
}

type byteCounter struct {
	w     io.Writer
	count int64
}

var _ io.Writer = (*byteCounter)(nil)

func (c *byteCounter) Write(p []byte) (int, error) {
	nbytes, err := c.w.Write(p)
	(*c).count += int64(nbytes)
	return nbytes, err
}
