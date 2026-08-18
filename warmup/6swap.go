package main

import "fmt"

func main() {
	a, b := 10, 30
	total := a + b
	a = total - a
	b = total - b
	fmt.Println(a, b)
}
