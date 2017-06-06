// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineToFiles := make(map[string]map[string]bool) // map[string]boolをsetとして利用
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineToFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, lineToFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s", n, line)
			var fileNames []string
			for fileName := range lineToFiles[line] {
				fileNames = append(fileNames, fileName)
			}
			fmt.Println("\tfound in " + strings.Join(fileNames, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineToFiles map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileNameSet, ok := lineToFiles[input.Text()]
		if !ok {
			fileNameSet = make(map[string]bool)
			lineToFiles[input.Text()] = fileNameSet
		}
		if !fileNameSet[f.Name()] {
			fileNameSet[f.Name()] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
