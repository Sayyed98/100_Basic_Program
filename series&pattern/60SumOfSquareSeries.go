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
		log.Println("error in converting input ", err)
		return
	}

	totalOfSqure := 0
	for i := 1; i <= num; i++ {
		totalOfSqure += i * i
	}

	fmt.Println("total of square series ", totalOfSqure)

	fmt.Println("sum ", Optimised(num))
}

func Optimised(num int) int {
	total := num * (num + 1) * (2*num + 1) / 6
	return total
}
