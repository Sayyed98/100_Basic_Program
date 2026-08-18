package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	total := 0
	for i := 1; i <= num; i = i + 2 {
		total += i

	}
	fmt.Println("total ", total)
}
