package main

import "fmt"

func main() {
	var arms int
	fmt.Scan(&arms)

	temp := arms
	total := 0

	length := 0
	for temp > 0 {
		length++
		temp /= 10
	}

	temp = arms
	for temp > 0 {
		rem := temp % 10
		total = total + PowerOf(rem, length)
		temp /= 10
	}
	if total == arms {
		fmt.Println("arms number")
	} else {
		fmt.Println("not arms")
	}
}

func PowerOf(num int, lenght int) int {
	total := 1
	for i := 0; i < lenght; i++ {
		total = total * num
	}

	return total
}
