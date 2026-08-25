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
	num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		log.Println("error in reading input ", err)
		return
	}

	for i := 1; i <= num; i++ {
		for j := num; j >= i; j-- {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

// inverted right triangle star pattern
