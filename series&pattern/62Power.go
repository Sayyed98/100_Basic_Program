package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println("x power of y is ", Power(x, y))
}

func Power(x, y int) int {
	total := 1
	for i := 1; i <= y; i++ {
		total *= x
	}
	return total
}
