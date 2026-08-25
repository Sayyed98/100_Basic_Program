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
		log.Println("error in reading input", err)
		return
	}
	num, err := strconv.Atoi(strings.TrimSpace(numString))
	if err != nil {
		log.Println("error in converting numstring to integer ", err)
		return
	}

	for i := 0; i < num; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(Pascal(i, j), " ")
		}
		fmt.Println()
	}
}
func Pascal(i, j int) int {
	if j == 0 || j == i {
		return 1
	}
	//fmt.Print("t ", i-1, " ", j-1, " ")
	return Pascal(i-1, j-1) + Pascal(i-1, j)
}
