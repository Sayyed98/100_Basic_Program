package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	temp := number
	fmt.Println("temp", temp)
	min := temp % 10
	for temp > 0 {
		rem := temp % 10
		fmt.Println("rem", rem)
		if rem < min {
			min = rem
		}
		temp = temp / 10
	}

	fmt.Println("smallest digit", min)
}
