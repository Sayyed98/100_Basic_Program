package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	max := 0
	if a > b {
		max = a
	} else {
		max = b
	}

	highest := 0
	for i := 2; i <= max; i++ {
		if a%i == 0 && b%i == 0 {
			highest = i
		}
	}

	fmt.Println("gcd or hcf ", highest)
}
