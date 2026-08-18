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

	fmt.Print("enter the base(m) number")
	m, err := readInt(reader)
	if err != nil {
		return
	}

	fmt.Print("enter the terms(n)")
	n, err := readInt(reader)
	if err != nil {
		return
	}

	for i := 1; i <= n; i++ {
		fmt.Println(i * m)
	}
}

func readInt(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Println("error in readng input ", err)
		return 0, err
	}

	return strconv.Atoi(strings.TrimSpace(input))
}
