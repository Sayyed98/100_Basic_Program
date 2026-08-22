package main

import "fmt"

func main() {
	var upto int
	fmt.Scan(&upto)

	for i := 1; i <= upto; i++ {
		if Armstrong(i) {
			fmt.Println(i)
		}
	}
}

func Armstrong(num int) bool {

	length := 0
	temp := num
	for temp > 0 {
		length++
		temp /= 10
	}

	temp = num
	total := 0
	for temp > 0 {
		rem := temp % 10
		total = total + PowerOff(rem, length)
		temp /= 10
	}

	return total == num
}

func PowerOff(num, length int) int {

	total := 1
	for i := 0; i < length; i++ {
		total = total * num
	}

	return total
}
