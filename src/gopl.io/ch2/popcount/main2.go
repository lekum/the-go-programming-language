package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	res := 0
	for i := uint(0); i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])

	}
	return res
}

func main() {
	for x := uint64(0); x < 32; x++ {
		fmt.Println(PopCount(x))
	}
}
