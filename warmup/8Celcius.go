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
	var celsius float64
	var fahrenheit float64

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter the celsius: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Print("error in reading celsius")
	}

	input = strings.TrimSpace(input)
	celsius, _ = strconv.ParseFloat(input, 64)
	fahrenheit = (celsius * 9 / 5) + 32

	fmt.Printf("%.2fC= %.2fF \n", celsius, fahrenheit)

}
