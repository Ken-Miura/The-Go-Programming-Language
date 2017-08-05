// Copyright 2017 Ken Miura
package archive

import "os"

type Item struct {
	Info     os.FileInfo
	Contents []byte
}
