package main

import "fmt"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println("Monday is", Monday)
	fmt.Println("Friday is", Friday)
}
