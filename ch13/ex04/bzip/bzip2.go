// Copyright 2017 Ken Miura
//!+

// Package bzip provides a writer that uses bzip2 compression (bzip.org).
package bzip

import (
	"fmt"
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	sync.Mutex
	w        io.Writer // underlying output stream
	bzip2    *exec.Cmd
	bzip2In  io.WriteCloser
	bzip2Out io.ReadCloser
	outbuf   []byte
}

// NewWriter returns a writer for bzip2-compressed streams.
func NewWriter(out io.Writer) io.WriteCloser {
	w := &writer{w: out}
	cmd := exec.Command("bzip2")
	w.bzip2 = cmd

	cmdIn, err := cmd.StdinPipe()
	if err != nil {
		panic(fmt.Sprintf("cannot get bzip2 stdin: %v", err))
	}
	w.bzip2In = cmdIn

	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		panic(fmt.Sprintf("cannot get bzip2 stdout: %v", err))
	}
	w.bzip2Out = cmdOut

	w.outbuf = make([]byte, 64*1024)

	if err := w.bzip2.Start(); err != nil {
		panic(fmt.Sprintf("cannot start bzip2 process: %v", err))
	}
	return w
}

//!-

//!+write
func (w *writer) Write(data []byte) (int, error) {
	w.Lock()
	defer w.Unlock()

	var total int // uncompressed bytes written
	for len(data) > 0 {
		n, err := w.bzip2In.Write(data)
		if err != nil {
			return total + n, err
		}
		total += n
		data = data[total:]
	}
	return total, nil
}

//!-write

//!+close
// Close flushes the compressed data and closes the stream.
// It does not close the underlying io.Writer.
func (w *writer) Close() error {
	w.Lock()
	defer w.Unlock()

	err := w.bzip2In.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(w.w, w.bzip2Out)
	if err != nil {
		return err
	}

	w.bzip2.Wait()
	return nil
}

//!-close
