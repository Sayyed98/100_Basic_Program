package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	temp := num
	sum := 0
	for temp > 0 {
		digit := temp % 10
		sum += factorial(digit)
		temp = temp / 10
	}

	if num == sum {
		fmt.Println("Strong number")
	} else {
		fmt.Println("Not strong number")
	}
}
func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)

}
