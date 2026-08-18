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
	fmt.Print("enter the number ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	number, _ := strconv.Atoi(input)

	total := 0
	for i := 1; i <= number; i++ {
		total += i
	}
	fmt.Println(total)
}
