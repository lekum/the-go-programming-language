package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func main() {
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}
