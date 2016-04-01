package main

import "fmt"

func mapequals(x, y map[string]string) bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok := y[xk]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func main() {
	x := map[string]string{
		"Jose Luis": "Álvarez",
		"María":     "Jiménez",
	}

	y := map[string]string{
		"Jose Luis": "Álvarez",
	}

	z := map[string]string{
		"Jose Luis": "Álvarez",
		"María":     "Jiménez",
	}

	fmt.Printf("x == y -> %t\n", mapequals(x, y))
	fmt.Printf("y == z -> %t\n", mapequals(y, z))
	fmt.Printf("x == z -> %t\n", mapequals(x, z))
	fmt.Printf("z == x -> %t\n", mapequals(z, x))
	fmt.Printf("y == x -> %t\n", mapequals(y, x))
}
