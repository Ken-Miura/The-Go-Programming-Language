// Copyright 2017 Ken Miura
package zip

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat("zip", 0, "PK", extract)
}

func extract(fileName string) ([]archive.Item, error) {
	r, err := zip.OpenReader(fileName)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var items []archive.Item
	for _, f := range r.File {
		var item archive.Item
		item.Info = f.FileInfo()
		item.Contents, err = getContents(f)
		if err != nil {
			return nil, fmt.Errorf("failed to extract %s: %v", f.Name, err)
		}
		items = append(items, item)
	}
	return items, nil
}

func getContents(f *zip.File) ([]byte, error) {
	rc, err := f.Open()
	if err != nil {
		return nil, err
	}
	defer rc.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, rc)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
