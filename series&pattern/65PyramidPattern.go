package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		for k := num; k > i; k-- {
			fmt.Print(" ")
		}
		for j := 1; j <= (2*i - 1); j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// pyramid pattern 