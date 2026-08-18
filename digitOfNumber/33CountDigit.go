package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter the number")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("err", err)
		return
	}

	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Println("error ", err)
		return
	}
	if num == 0 {
		fmt.Println("count is ", 1)
		return
	}
	if num < 0 {
		num = -num
	}
	count := 0
	temp := num
	for temp > 0 {

		count++
		temp = temp / 10
	}
	fmt.Println("digit count", count)
}
