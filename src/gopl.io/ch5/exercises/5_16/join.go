package main

import "fmt"

func Join(sep string, strs ...string) string {
	var res string
	res = strs[0]
	for _, str := range strs[1:] {
		res = res + sep + str
	}
	return res

}

func main() {
	str1 := "Hello"
	str2 := "World"
	fmt.Println(Join(" ", str1, str2))
}
