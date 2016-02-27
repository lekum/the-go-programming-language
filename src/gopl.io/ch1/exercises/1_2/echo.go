package main

import (
	"fmt"
	"os"
)

func main() {
	for i, value := range os.Args[1:] {
		fmt.Printf("%d -> %s\n", i+1, value)
	}
}
