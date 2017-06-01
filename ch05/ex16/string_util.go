// Copyright 2017 Ken Mirua
package ex16

import "strings"

func Join(sep string, strs ...string) string {
	return strings.Join(strs[:], sep)
}
