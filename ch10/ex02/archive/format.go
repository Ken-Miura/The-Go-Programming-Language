// Copyright 2017 Ken Miura
package archive

import (
	"bufio"
	"errors"
	"io"
	"os"
)

var ErrFormat = errors.New("archive: unknown format")

// formatは、アーカイブメソッド名、magicまでのオフセット、アーカイブメソッドを識別するバイナリ (magic)、データ抽出方法を含む。
type format struct {
	name    string
	offset  int
	magic   string
	extract func(string) ([]Item, error)
}

var formats []format

func RegisterFormat(name string, offset int, magic string, extract func(string) ([]Item, error)) {
	formats = append(formats, format{name, offset, magic, extract})
}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) format {
	for _, f := range formats {
		b, err := r.Peek(f.offset + len(f.magic))
		if err == nil && match(f.magic, b[f.offset:]) {
			return f
		}
	}
	return format{}
}

// Extract extracts an archive that has been archived in a registered format.
// The string returned is the format name used during format registration.
// Format registration is typically done by an init function in the codec-
// specific package.
func Extract(fileName string) ([]Item, string, error) {
	format, err := detectFormat(fileName)
	if err != nil {
		return nil, "", err
	}
	if format.extract == nil {
		return nil, "", ErrFormat
	}
	items, err := format.extract(fileName)
	return items, format.name, err
}

func detectFormat(fileName string) (format, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return format{}, err
	}
	defer f.Close()
	rr := asReader(f)
	return sniff(rr), nil
}
