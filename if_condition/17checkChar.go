package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a character for checking its alphabet, digit or symbol")
	scanner.Scan()

	input := scanner.Text()
	if input >= "a" && input <= "z" || input >= "A" && input <= "Z" {
		fmt.Println("alphabet")
	} else if input >= "1" && input <= "9" {
		fmt.Println("digit")
	} else {
		fmt.Println("special character")
	}
}
