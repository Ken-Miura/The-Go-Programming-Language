// Copyright 2017 Ken Miura
package ex05

import "io"

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n, 0}
}

type limitReader struct {
	r    io.Reader
	n    int64
	read int64
}

var _ io.Reader = (*limitReader)(nil)

func (lr *limitReader) Read(p []byte) (int, error) {
	nbytes, err := lr.r.Read(p)
	if (lr.read + int64(nbytes)) >= lr.n {
		return int(lr.n - lr.read), io.EOF
	}
	lr.read += int64(nbytes)
	return nbytes, err
}
