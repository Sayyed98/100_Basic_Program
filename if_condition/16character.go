package main

import "fmt"

func main() {
	var char string
	fmt.Scan(&char)

	if char == "a" || char == "e" || char == "i" || char == "o" || char == "u" {
		fmt.Println("vowel", char)
	} else {
		fmt.Println("consonant", char)
	}
}
