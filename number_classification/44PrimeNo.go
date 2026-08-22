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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	number, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		log.Println("error in scanner text and converting", err)
		return
	}

	if number%2 == 0 || number%3 == 0 {
		fmt.Println("not prime number")
		return
	}
	isPrime := true
	for i := 5; i*i <= number; i += 6 {
		if number%i == 0 || number%(i+2) == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		fmt.Println("prime nuber")
	} else {
		fmt.Println("not prime number")
	}
}
