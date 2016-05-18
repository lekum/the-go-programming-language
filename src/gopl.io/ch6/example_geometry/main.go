package main

import (
	"fmt"
	"gopl.io/ch6/geometry"
)

func main() {
	x := geometry.Point{X: 1, Y: 2}
	y := geometry.Point{X: 3, Y: 4}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("Distance: %v\n", x.Distance(y))
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}
