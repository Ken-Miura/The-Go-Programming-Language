// Copyright 2017 Ken Miura
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("enter expression: ")
		input, _, err := in.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error: %v \n", err)
			return
		}
		env := make(Env, 0)
		expr, err := parseAndInquireVar(string(input), env)
		fmt.Printf("evaluation result: %g\n", expr.Eval(env))
	}
}
