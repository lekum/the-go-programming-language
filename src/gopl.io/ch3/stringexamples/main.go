package main

import "fmt"

func main() {
	s := "Feliz Año nuevo"
	fmt.Println(len(s))
	for ix, v := range s {
		fmt.Printf("%d -> %c\n", ix, v)
	}
}
