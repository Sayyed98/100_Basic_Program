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
	var min int
	for i := 0; i < len(num); i++ {
		scanner.Scan()
		num[i], _ = strconv.Atoi(scanner.Text())
	}

	min = num[0]
	for _, v := range num {
		if min > v {
			min = v
		}
	}
	fmt.Println("minimum number is ", min)
}
