package main

import "fmt"

func main() {
	var radius float64
	fmt.Scanln(&radius)

	circumference := 2 * 3.14 * radius
	fmt.Println("circumference", circumference)

	area := 3.14 * radius * radius
	fmt.Println("area", area)
}
