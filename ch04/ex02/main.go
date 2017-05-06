// Copyright 2017 Ken Mirua
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var algorithm = flag.String("a", "sha256", "sha algorithm: sha256, sha384 or sha512")

func main() {
	flag.Parse()
	args := flag.Args()
	if !(*algorithm == "sha256" || *algorithm == "sha384" || *algorithm == "sha512") {
		fmt.Println("Algorithm must be sha256, sha384 or sha512.")
		return
	}
	if len(args) != 1 {
		fmt.Println("usage: " + os.Args[0] + " [-a 'algorithm'] message you want to hash")
		fmt.Println("ex1. " + os.Args[0] + " x")
		fmt.Println("ex1. " + os.Args[0] + " -a 'sha512' y")
		return
	}

	if *algorithm == "sha256" {
		fmt.Printf("%x", sha256.Sum256([]byte(args[0])))
	} else if *algorithm == "sha384" {
		fmt.Printf("%x", sha512.Sum384([]byte(args[0])))
	} else if *algorithm == "sha512" {
		fmt.Printf("%x", sha512.Sum512([]byte(args[0])))
	} else {
		panic("This line must not be reached.\n")
	}
}
