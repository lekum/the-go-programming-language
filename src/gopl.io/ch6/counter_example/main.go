package main

import (
	"fmt"
	"gopl.io/ch6/counter"
)

func main() {
	c1 := counter.Counter{}
	c2 := counter.Counter{}
	c2.Increment()
	fmt.Println(c1.N())
	fmt.Println(c2.N())
	c2.Reset()
	fmt.Println(c2.N())
}
