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

	numString, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error in reading input", numString)
		return
	}

	num, err := strconv.Atoi(strings.TrimSpace(numString))
	if err != nil {
		log.Println("error in converting input", numString)
		return
	}

	a, b := 0, 1
	fmt.Println(a)
	fmt.Println(b)

	for i := 2; i <= num; i++ {
		a, b = b, a+b

		fmt.Println(b)
	}
	fmt.Println("fibon", FiboRec(num))
}

func FiboRec(num int) int {
	if num <= 1 {
		return num
	}
	return FiboRec(num-1) + FiboRec(num-2)
}
