package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	for i := 2; i <= num; i++ {
		if PrimeNo(i) {
			fmt.Println(i)
		}
	}

}

func PrimeNo(num int) bool {
	if num < 2 {
		fmt.Println("not a prime number")
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
