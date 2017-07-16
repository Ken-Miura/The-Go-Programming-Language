// Copyright 2017 Ken Miura
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	node, err := ConstructXmlNodeTree(os.Stdin)
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
		fmt.Fprintf(out, "%*s%s\n", depth*2, "", strings.TrimSpace(string(node)))
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
	var rootNode, currentNode, currentNodeParent Node = nil, nil, nil
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			return rootNode, nil
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			elem := Element{tok.Name, tok.Attr, nil}
			if rootNode == nil {
				rootNode = &elem
			}
			if currentNode != nil {
				element := currentNode.(*Element)
				element.Children = append(element.Children, &elem)
			}
			currentNodeParent = currentNode
			currentNode = &elem
		case xml.EndElement:
			currentNode = currentNodeParent
		case xml.CharData:
			element, ok := currentNode.(*Element)
			if ok {
				element.Children = append(element.Children, CharData(tok))
			}
		default:
			// その他のトークンは無視
		}
	}
	panic("This line must not be reached.")
}
