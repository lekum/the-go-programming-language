package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Printf("x=%b\n", x)
	var err error
	x, err = strconv.Atoi("123")
	fmt.Println(x)
	var z int64
	z, err = strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	if err != nil {
		return
	}
	fmt.Println(z)
}
