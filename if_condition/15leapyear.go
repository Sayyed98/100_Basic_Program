package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num%4 == 0 && num%100 != 0 || num%4 == 0 {
		fmt.Println("leap yeear")
	} else {
		fmt.Println("not leap year")
	}
	fmt.Println(num)

}
