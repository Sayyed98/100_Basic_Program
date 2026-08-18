package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var num int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("please enter the number")
	scanner.Scan()

	val := scanner.Text()

	num, err := strconv.Atoi(val)
	if err != nil {
		log.Println("error is : ", err)
		return
	}

	if num > 0 {
		fmt.Println("number is positive")
	} else if num == 0 {
		fmt.Println("number is zero", num)
	} else {
		fmt.Println("number is negative")
	}

}
