// Copyright 2017 Ken Miura
package ex03

import (
	"strings"
)

var echoTests = []string{"arg0", "arg1", "arg2", "arg3", "arg4", "arg5", "arg6", "arg7", "arg8", "arg9"}

func EchoInefficient() {
	s, sep := "", ""
	for _, arg := range echoTests {
		s += sep + arg
		sep = " "
	}
}

func EchoEfficient() {
	strings.Join(echoTests, " ")
}
