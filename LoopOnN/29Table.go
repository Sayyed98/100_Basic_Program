package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	total := 0
	for i := 1; i <= 10; i++ {
		total = num * i
		fmt.Println(total)
	}

}
n
