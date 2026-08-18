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
	fmt.Println("enter the person age")
	scanner.Scan()
	input := scanner.Text()
	input = strings.TrimSpace(input)

	age, _ := strconv.Atoi(input)

	if age >= 18 {
		fmt.Println("person is eligible to vote")
	} else {
		fmt.Println("not eligible")
	}
}
