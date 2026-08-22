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

	read, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error in reading input")
		return
	}

	number, err := strconv.Atoi(strings.TrimSpace(read))
	if err != nil {
		log.Println("err in converting string to integer")
		return
	}

	if number <= 0 {
		fmt.Println("enter a positive number")
		return
	}

	temp := number
	total := 0
	for temp > 0 {
		digit := temp % 10
		total = total + digit
		temp /= 10
	}

	if number%total == 0 {
		fmt.Println("nived number")
	} else {
		fmt.Println("not nived number")
	}
}
