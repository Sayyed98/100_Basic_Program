package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var num [3]int

	for i := 0; i < len(num); i++ {
		scanner.Scan()
		num[i], _ = strconv.Atoi(scanner.Text())
	}

	max := 0
	for i := 0; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}
	}

	fmt.Println("largest number is ", max)
}
