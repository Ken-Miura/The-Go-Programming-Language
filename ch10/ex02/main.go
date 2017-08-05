// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"os"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive"
	_ "github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive/tar"
	_ "github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive/zip"
)

func main() {
	test := "tar_test.tar"
	// test = "zip_test.zip"
	items, method, err := archive.Extract(test)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(method)
	for _, item := range items {
		fmt.Println(item.Info.Name())
	}
}
