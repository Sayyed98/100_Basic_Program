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
	fmt.Println("enter the second value: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error in reading input ", err)
		return
	}
	input = strings.TrimSpace(input)
	second, err := strconv.Atoi(input)
	if err != nil {
		log.Println("please provide correct input", second)
		return
	}

	hours := second / 3600
	minute := (second % 3600) / 60
	remainingSecond := second % 60
	fmt.Printf("hours: %d minutes %d second %d \n", hours, minute, remainingSecond)
}
