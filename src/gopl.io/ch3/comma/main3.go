package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		buf.WriteRune(v)
		if i%3 == 0 && i < len(s)-1 {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567"))
}
