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
	fmt.Println("enter the number")

	strNumber, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error in reading the input", err)
		return
	}

	num, err := strconv.Atoi(strings.TrimSpace(strNumber))
	if err != nil {
		log.Println("error in converting string number to number")
		return
	}
	count := 2
	fmt.Println(num, count)
	for num > 0 {
		if Prime(count) {
			fmt.Println(count)
			num--
		}
		count++

	}

}
func Prime(num int) bool {
	if num < 2 {
		fmt.Println("not prime number")
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
