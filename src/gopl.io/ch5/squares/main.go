package main

import "fmt"

func squares(x int) func() int {
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares(0)
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
