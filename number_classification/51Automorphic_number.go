package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	square := num * num

	// length
	temp := num
	length := 0

	for temp > 0 {
		_ = temp % 10
		length++
		temp = temp / 10
	}

	fmt.Println("lenght and squre", length, square)
	temp = square
	place := 1
	total := 0
	isMorphic := false
	for temp > 0 {
		digit := temp % 10
		fmt.Println("digit", digit)
		total = total + digit*place
		fmt.Println("total", total)

		place = place * 10
		if total == num {
			isMorphic = true
			break
		}
		temp /= 10

	}
	if isMorphic {
		fmt.Println("morphic")
	} else {
		fmt.Println("not morhic")
	}
}

/**

optimised approach
func main() {
	var num int
	fmt.Scan(&num)

	square := num * num

	// Count digits
	temp := num
	length := 0

	for temp > 0 {
		length++
		temp /= 10
	}

	// Create 10^length
	pow := 1
	for i := 0; i < length; i++ {
		pow *= 10
	}

	if square%pow == num {
		fmt.Println("Automorphic")
	} else {
		fmt.Println("Not Automorphic")
	}
}
**/
