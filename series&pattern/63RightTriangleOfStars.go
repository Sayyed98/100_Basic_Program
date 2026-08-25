package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// right triangle star pattern
