// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[:] {
		fmt.Println(strconv.Itoa(i) + ", " + arg)
	}
}
