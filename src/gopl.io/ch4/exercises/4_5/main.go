package main

import "fmt"

func uniq(list []string) []string {
	var end = len(list)
	for i := 1; i < end; i++ {
		if list[i-1] == list[i] {
			copy(list[i-1:], list[i:])
			end--
			i--
		}

	}
	return list[:end]
}

func main() {
	s := []string{"one", "two", "two", "two", "three", "three", "four",
		"five", "five", "six"}
	fmt.Println(s)
	fmt.Println(uniq(s))
}
