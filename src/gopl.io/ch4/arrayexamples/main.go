package main

import "fmt"

func main() {
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)
	r := [...]string{1: "hello", 2: "world"}
	fmt.Printf("%v, type %T\n", r, r)
}
