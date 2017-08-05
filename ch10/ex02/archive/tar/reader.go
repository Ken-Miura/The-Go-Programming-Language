// Copyright 2017 Ken Miura
package tar

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"

	"bytes"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive"
)

func init() {
	archive.RegisterFormat("tar", 0x101, "ustar", extract)
}

func extract(fileName string) ([]archive.Item, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tr := tar.NewReader(f)

	var items []archive.Item
	// Iterate through the files in the archive.
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		var item archive.Item
		item.Info = hdr.FileInfo()
		item.Contents, err = getContents(tr)
		if err != nil {
			return nil, fmt.Errorf("failed to extract %s: %v", hdr.Name, err)
		}
		items = append(items, item)
	}
	return items, nil
}

func getContents(tr *tar.Reader) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, tr)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
