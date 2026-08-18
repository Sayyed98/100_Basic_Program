package main

import "fmt"

func main() {
	a, b := 20, 40
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}
