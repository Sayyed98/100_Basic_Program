package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	var total int = 1
	// you need to assign total 1 otherwise it will become 0, 0 multiply by n is zero
	for i := 1; i <= num; i++ {
		total *= i
	}

	fmt.Println("total ", total)
}
