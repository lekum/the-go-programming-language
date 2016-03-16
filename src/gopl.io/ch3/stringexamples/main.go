package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Feliz Año nuevo"
	fmt.Println(s)
	fmt.Println("Length:", len(s))
	fmt.Println("Rule count:", utf8.RuneCountInString(s))
	for ix, v := range s {
		fmt.Printf("%d -> %q, %d\n", ix, v, v)
	}
	t := "Feliz \xe4\xb8\x96\xe7\x95\x8c"
	fmt.Println(t)
	fmt.Println(len(t))
	fmt.Println("Rule count:", utf8.RuneCountInString(t))
	for ix, v := range t {
		fmt.Printf("%d -> %q, %d\n", ix, v, v)
	}
}
