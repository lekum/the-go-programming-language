package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		lines[input.Text()]++
	}

	for key, value := range lines {
		fmt.Println("The line", key, "is repeated", value, "times")
	}
}
