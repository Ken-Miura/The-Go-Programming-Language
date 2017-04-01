// Copyright 2017 Ken Mirua
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineToFiles := make(map[string][]string)
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
			fmt.Println("\tfound in " + strings.Join(lineToFiles[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineToFiles map[string][]string) {
	input := bufio.NewScanner(f)
Counts:
	for input.Scan() {
		counts[input.Text()]++
		for _, file := range lineToFiles[input.Text()] {
			if file == f.Name() {
				continue Counts
			}
		}
		lineToFiles[input.Text()] = append(lineToFiles[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
