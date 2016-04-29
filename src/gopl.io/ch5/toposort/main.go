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
}

func main() {
	fmt.Println("topoSort:")
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	fmt.Println("topoSort2:")
	for i, course := range topoSort2(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

}

// depth first search
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
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

func topoSort2(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)
	visitAll = func(items map[string]bool) {
		for item, _ := range items {
			if !seen[item] {
				seen[item] = true
				map2 := make(map[string]bool)
				for _, key2 := range m[item] {
					map2[key2] = true
				}
				visitAll(map2)
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
