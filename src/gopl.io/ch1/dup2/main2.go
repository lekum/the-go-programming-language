package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineCount struct {
	count int
	files []string
}

func main() {
	counts := make(map[string]*lineCount)
	if len(os.Args) == 1 {
		countlines(os.Stdin, "stdin", counts)
	} else {
		files := os.Args[1:]

		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 error: %v\n", err)
			}
			countlines(f, file, counts)
			f.Close()
		}
		for line, c := range counts {
			if c.count > 1 {
				fmt.Printf("%d\t%s\t%v\n", c.count, line, c.files)
			}
		}
	}
}

func countlines(f *os.File, filename string, c map[string]*lineCount) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		cline, ok := c[line]
		if !ok {
			// Creating a new lineCount and referencing it from the mapping
			cline = new(lineCount)
			c[line] = cline
		}
		cline.count++
		cline.files = append(cline.files, filename)
	}
}
