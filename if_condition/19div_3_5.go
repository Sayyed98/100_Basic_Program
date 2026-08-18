package main

import "fmt"

func main() {
	var num int
	fmt.Scanln(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("number is divisible")
	} else {
		fmt.Println("not divisible")
	}
}
