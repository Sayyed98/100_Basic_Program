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
	fmt.Println("please enter the marks")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error in reading input ", err)
		return
	}
	input = strings.TrimSpace(input)
	marks, err := strconv.Atoi(input)
	if err != nil {
		log.Println("please provide correct input", err)
		return
	}
	fmt.Println(marks, "makrs")
	if marks >= 80 {
		fmt.Println("grade A")
	} else if marks >= 60 && marks < 80 {
		fmt.Println("grade B")
	} else if marks >= 50 && marks < 60 {
		fmt.Println("grade c")
	} else if marks >= 33 && marks < 50 {
		fmt.Println("grade d")
	} else {
		fmt.Println("fail")
	}

}
