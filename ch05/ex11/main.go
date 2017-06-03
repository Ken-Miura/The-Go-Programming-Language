// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},

	"linear algebra": {"calculus"}, // 循環依存の追加
}

func main() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				var dependencyCheck func(strs []string)
				dependencyCheck = func(strs []string) {
					for _, s := range strs {
						if s == item {
							panic(s + " has cyclic dependency")
						}
						dependencyCheck(m[s])
					}
				}
				dependencyCheck(m[item])
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)
	return order
}
