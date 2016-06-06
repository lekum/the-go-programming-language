package main

import (
	"fmt"
	"gopl.io/ch6/geometry"
	"image/color"
)

type ColoredPoint struct {
	*geometry.Point
	Color color.RGBA
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{&geometry.Point{1, 1}, red}
	q := ColoredPoint{&geometry.Point{5, 4}, blue}
	fmt.Println(p.Distance(*q.Point))
	q.Point = p.Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point)
}
