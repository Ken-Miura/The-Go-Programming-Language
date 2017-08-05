// Copyright 2017 Ken Miura
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive"
	_ "github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive/tar"
	_ "github.com/Ken-Miura/The-Go-Programming-Language/ch10/ex02/archive/zip"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " 'archive file you want to extract'")
		fmt.Println("ex: " + os.Args[0] + " tar_test.tar")
		os.Exit(0)
	}
	items, method, err := archive.Extract(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("archive method: " + method)
	for _, item := range items {
		err = createFile(item.Info.Name(), item.Contents)
		if err != nil {
			fmt.Printf("failed to create file %s: %v\n", item.Info.Name(), err)
		}
	}
}

func createFile(fileName string, contents []byte) (err error) {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	buf := bytes.NewBuffer(contents)
	_, err = io.Copy(f, buf)
	if err != nil {
		return err
	}
	return nil
}
