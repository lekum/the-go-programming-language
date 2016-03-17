package main

import (
	"fmt"
	"os"
	"strings"
)

func comma(s string) string {
	var fract string
	if p := strings.LastIndex(s, "."); p >= 0 {
		fract = s[p:]
		s = s[:p]
	}
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:] + fract
}

func main() {
	s := os.Args[1]
	fmt.Println(comma(s))
}
