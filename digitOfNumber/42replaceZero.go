package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	temp := num
	total := 0
	place := 1
	for temp > 0 {
		digit := temp % 10
		if digit == 0 {
			digit = 5
		}

		total = digit*place + total
		place = place * 10
		temp = temp / 10
	}

	fmt.Println(total)
}
