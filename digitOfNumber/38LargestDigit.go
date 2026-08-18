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

	read, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error in reading input", err)
		return
	}

	num, err := strconv.Atoi(strings.TrimSpace(read))
	if err != nil {
		log.Println("error in converting ")
	}
	temp := num
	fmt.Println(temp)
	max := 0
	for temp > 0 {
		rem := temp % 10
		if rem > max {
			max = rem
		}
		temp = temp / 10
	}

	fmt.Println("largest digit", max)

}
