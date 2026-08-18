package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	temp := num
	firstDigit := 0
	lastDigit := 0
	isFirst := true
	for temp > 0 {
		rem := temp % 10
		if isFirst {
			lastDigit = rem
			isFirst = false
		}
		if temp < 10 {
			firstDigit = rem
		}

		temp = temp / 10
	}
	fmt.Println("total", firstDigit+lastDigit)
}
