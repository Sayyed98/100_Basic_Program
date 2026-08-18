package main

import (
	"fmt"
	"log"
)

func main() {
	var numbers int
	_, err := fmt.Scan(&numbers)
	if err != nil {
		log.Println("error  is : ", err)
		return
	}

	if numbers%2 != 0 {
		fmt.Println("odd number")
	} else {
		fmt.Println("even number")
	}
}
