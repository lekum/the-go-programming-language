package main

import (
	"fmt"
	"gopl.io/ch6/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	distance := geometry.Point.Distance // method expression
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*geometry.Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}
