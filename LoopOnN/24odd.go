package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter the number Upto")
	scanner.Scan()

	input := scanner.Text()
	input = strings.TrimSpace(input)
	number, _ := strconv.Atoi(input)

	for i := 1; i <= number; i = i + 2 {
		fmt.Println(i)
	}
}
