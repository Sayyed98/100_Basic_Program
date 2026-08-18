package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	count := 2

	for num > 0 {
		if Primes(count) {
			fmt.Println(count)
			num--
		}

		count++
	}
}

func Primes(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
