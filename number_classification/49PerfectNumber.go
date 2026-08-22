package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum = sum + i
		}
	}

	if num == sum {
		fmt.Println("perfect number")
	} else {
		fmt.Println("not a perfect number")
	}
}

// optimised one
func PerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}

	sum := 1

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			fmt.Println("sum outside of if conditin", sum, i)
			if i != num/i {
				sum += num / i
				fmt.Println("sum inside", sum, i)
			}
		}
	}

	return sum == num
}
