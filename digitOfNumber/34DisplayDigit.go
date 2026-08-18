package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	temp := number
	arr := []int{}
	for temp > 0 {
		rem := temp % 10
		arr = append(arr, rem)
		temp = temp / 10
	}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[len(arr)-i-1])
	}
	fmt.Println("printing digit of number line by line")

	digitRec(435)
}

// other way using recursion
func digitRec(num int) {
	if num == 0 {
		return
	}
	digitRec(num / 10)
	fmt.Println(num % 10)
}
