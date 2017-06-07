// Copyright 2017 Ken Miura
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type element struct {
	name       string
	attributes []attribute
}

type attribute struct {
	key   string
	value string
}

func main() {
	elements := parseArgs()
	dec := xml.NewDecoder(os.Stdin)
	var stack []element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			var attrs []attribute
			for _, v := range tok.Attr {
				attrs = append(attrs, attribute{v.Name.Local, v.Value})
			}
			stack = append(stack, element{tok.Name.Local, attrs}) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, elements) {
				for _, elem := range stack {
					fmt.Printf("%s ", elem.name)
					for _, attr := range elem.attributes {
						fmt.Printf(`%s="%s" `, attr.key, attr.value)
					}
				}
				fmt.Printf(": %s\n", tok)
			}
		}
	}
}

func parseArgs() []element {
	var elements []element
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "=") {
			if len(elements) == 0 {
				panic("invalid args: " + strings.Join(os.Args[1:], " ")) // 要素がないのに属性が来るのは、入力値がおかしいのでpanicで処理
			}
			keyAndValue := strings.Split(arg, "=")
			key := keyAndValue[0]
			value := keyAndValue[1]
			attrs := elements[len(elements)-1].attributes
			attrs = append(attrs, attribute{key, value})
			elements[len(elements)-1].attributes = attrs
		} else {
			elem := element{arg, nil}
			elements = append(elements, elem)
		}
	}
	return elements
}

func containsAll(x, y []element) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if equalElement(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func equalElement(x, y element) bool {
	if x.name != y.name {
		return false
	}
	if len(x.attributes) != len(y.attributes) {
		return false
	}
	for i := range x.attributes {
		if x.attributes[i] != y.attributes[i] {
			return false
		}
	}
	return true
}
