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
	fmt.Println("enter the number upto n")
	scanner.Scan()

	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Println("err ", err)
		return
	}

	for i := 1; i <= number; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i)
		}
	}
}
