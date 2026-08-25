package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	a, b := 0, 1
	c := 0
	sum := 0
	for i := 2; i <= num; i++ {

		c = a + b
		a = b
		b = c
		sum = sum + a
	}
	fmt.Println("sum of fibo ", sum)
}
