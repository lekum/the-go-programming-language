package main

import (
	"fmt"
	"os"
)

func main() {
	var line, sep = "", ""
	for _, arg := range os.Args[1:] {
		line += sep + arg
		sep = " "
	}
	fmt.Println(line)
}
