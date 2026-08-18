package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("enter the number ")
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Println("error in conversion ", err)
		return
	}

	temp := num

	if temp == 0 {
		fmt.Println("input is zero")
		return
	}
	if temp < 0 {
		temp = -temp
	}

	total := 1
	for temp > 0 {
		digit := temp % 10
		total = total * digit
		temp = temp / 10
	}
	fmt.Println("total is ", total)
}
