package main

import "fmt"

func appendInt(x []int, i int) (y []int) {
	var z []int
	xlen := len(x)
	zlen := xlen + 1
	xcap := cap(x)
	if zlen < xcap {
		// Enough space in x
		z = x[:zlen]
	} else {
		// Grow x by doubling
		zcap := zlen
		if zcap < 2*xlen {
			zcap = 2 * xlen
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[xlen] = i
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
