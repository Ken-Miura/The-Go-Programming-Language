// Copyright 2017 Ken Miura
package main

import (
	"fmt"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	trials := 1000
	for i := 0; i < trials; i++ {
		validateToposortResult(prereqs, topoSort(prereqs))
	}
	fmt.Printf("\nvalidation (%d times) passed\n", trials)
}

func validateToposortResult(prereqs map[string]map[string]bool, nodes []string) {
NEXT_NODE:
	for i, v := range nodes {
		dependensySet := prereqs[v]
		subResult := nodes[:i]
		for dependency := range dependensySet {
			for _, s := range subResult {
				if s == dependency {
					continue NEXT_NODE
				}
			}
			panic("dependency is not satisfied")
		}
	}
}

func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order
}
