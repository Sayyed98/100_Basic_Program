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
	fmt.Print("enter the number  ")

	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	number, _ := strconv.Atoi(input)

	for i := number; i >= 1; i-- {
		fmt.Println(i)
	}
}
