// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func main() {
	breadthFirst(crawlDir, os.Args[1:])
}

func crawlDir(dir string) []string {
	fmt.Println(dir)
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var dirs []string
	for _, item := range items {
		if item.IsDir() {
			dirs = append(dirs, dir+string(filepath.Separator)+item.Name())
		} else {
			fmt.Println(dir + string(filepath.Separator) + item.Name())
		}
	}
	return dirs
}
