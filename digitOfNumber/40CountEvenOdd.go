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
	fmt.Println("enter the input")
	scanner.Scan()

	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Println("error in converting number", err)
		return
	}

	temp := num
	evenCount := 0
	oddCount := 0
	for temp > 0 {
		rem := temp % 10
		if rem%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
		temp = temp / 10
	}

	fmt.Println("evenCount", evenCount, "odd Count", oddCount)
}
