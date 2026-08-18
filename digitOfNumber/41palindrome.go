package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the number")
	number, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(number))

	temp := num
	total := 0
	for temp > 0 {
		rem := temp % 10
		total = total*10 + rem
		temp = temp / 10
	}

	if total == num {
		fmt.Println("palindrome ", total)
	} else {
		fmt.Println("Not palindrome ", total)
	}
}
