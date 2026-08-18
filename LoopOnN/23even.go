package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	for i := 2; i <= num; i = i + 2 {
		fmt.Println(i)
	}
}
