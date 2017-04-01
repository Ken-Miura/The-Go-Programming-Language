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
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

			f2, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			findLines(f2, lineToFiles)
			f2.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Println(line + ": found in " + strings.Join(lineToFiles[line], ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func findLines(f *os.File, lineToFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		for _, file := range lineToFiles[input.Text()] {
			if file == f.Name() {
				return
			}
		}
		lineToFiles[input.Text()] = append(lineToFiles[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}
