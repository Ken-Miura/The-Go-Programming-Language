// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"fmt"
	"os"
)

// TODO split functionを自作して正規表現 ([\\P{L}]+とか) でトークンを分ければもっときれいに単語の切り分けできそう？
func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: " + os.Args[0] + " 'text file in which you want to count words'")
		fmt.Println("ex. " + os.Args[0] + " alice.txt")
		return
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("failed to open file. error: %v", err)
		return
	}
	defer f.Close()

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)

	wordfreq := make(map[string]int)
	for input.Scan() {
		wordfreq[input.Text()]++
	}

	fmt.Println("word : number of times that the word appears")
	for k, v := range wordfreq {
		fmt.Printf("%s : %d\n", k, v)
	}
}
