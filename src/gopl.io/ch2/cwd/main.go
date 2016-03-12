package main

import (
	"fmt"
	"log"
	"os"
)

var cwd string

func main() {
	var err error
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error trying to get current working directory: %v\n", err)
	}
	fmt.Println("The current working directory is", cwd)

}
