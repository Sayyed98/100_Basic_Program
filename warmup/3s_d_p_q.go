package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanln(&a, &b)
	add := a + b
	fmt.Println(add)
	diff := a - b
	fmt.Println(diff)

	product := a * b
	fmt.Println(product)

	quotient := a / b
	fmt.Println(quotient)
}
