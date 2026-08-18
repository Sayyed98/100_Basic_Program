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
	fmt.Print("enter the number ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error is ", err)
		return
	}

	input = strings.TrimSpace(input)
	num, _ := strconv.Atoi(input)

	temp := num
	total := 0
	for temp > 0 {
		remainder := temp % 10
		total = total*10 + remainder
		temp = temp / 10
	}
	fmt.Println(total)
}
