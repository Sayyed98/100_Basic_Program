package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	var total int
	for i := 0; i <= num; i = i + 2 {
		total += i
	}
	fmt.Println("total ", total)
}
