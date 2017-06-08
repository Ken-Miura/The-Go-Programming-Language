// Copyright 2017 Ken Miura
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("books.xml")
	node, err := ConstructXmlNodeTree(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "constructing xml node tree: %v\n", err)
		return
	}
	printXmlNodeTree(os.Stdout, node, 0)
}

func printXmlNodeTree(out io.Writer, node Node, depth int) {
	switch node := node.(type) {
	case *Element:
		fmt.Printf("%*s<%s>\n", depth*2, "", string(node.Type.Local))
		for _, childNode := range node.Children {
			printXmlNodeTree(out, childNode, depth+1)
		}
	case CharData:
		fmt.Fprintf(out, "%*s%s\n", depth*2, "", string(node))
	default:
		panic("This line must not be reached.\n")
	}
}

type Node interface{} // *ElementまたはCharData

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func ConstructXmlNodeTree(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var root, currentNode, currentNodeParent Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := Element{tok.Name, tok.Attr, nil}
			if root != nil {
				parent, ok := currentNode.(*Element)
				if ok {
					parent.Children = append(parent.Children, &elem)
					currentNodeParent = parent
					currentNode = &elem
				}
			} else {
				root = &elem
				currentNode = root
			}
		case xml.EndElement:
			currentNode = currentNodeParent
			currentNodeParent = nil
		case xml.CharData:
			parent, ok := currentNode.(*Element)
			if ok {
				parent.Children = append(parent.Children, CharData(tok))
			}
		}
	}
	return root, nil
}
