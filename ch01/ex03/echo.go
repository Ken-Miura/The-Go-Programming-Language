// Copyright 2017 Ken Miura
package ex03

import (
	"fmt"
	"os"
	"strings"
)

func EchoInefficient() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func EchoEfficient() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
