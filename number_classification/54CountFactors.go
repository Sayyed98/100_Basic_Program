package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	count := 0

	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(count, i)
			count++
		}
	}

	fmt.Println("total count of factors of number", count)
}
