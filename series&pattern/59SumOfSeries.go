package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}

	fmt.Println("sum of the series ", total)
}
