package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	total := 1.0

	for i := 2; i <= num; i++ {
		div := 1.0 / float64(i)
		total += div
		fmt.Println("total", total)
	}

	fmt.Println("sum of series 1+1/n series ", total)
}
