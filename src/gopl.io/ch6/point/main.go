package main

import "fmt"

type Point struct{ X, Y float64 }

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}

func main() {
	path_1 := Path{
		Point{1, 0},
		Point{2, 1},
		Point{3, 1},
	}

	pivot := Point{5, 5}
	path_1.TranslateBy(pivot, true)
	fmt.Println(path_1)
	path_1.TranslateBy(pivot, false)
	fmt.Println(path_1)
}
