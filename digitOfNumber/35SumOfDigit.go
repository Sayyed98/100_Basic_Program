package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	temp := num

	if num == 0 {
		fmt.Println(0)
		return
	}

	if temp < 0 {
		temp = -temp
	}
	sum := 0
	for temp > 0 {
		digit := temp % 10
		sum = sum + digit
		temp = temp / 10
	}
	fmt.Println("sum of digit ", sum)
}
