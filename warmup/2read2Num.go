package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var err error
	_, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Errorf("error", err)
		return

	}

	// _, err = fmt.Scanf("b %d ", &b)
	// if err != nil {
	// 	fmt.Errorf("error ", err)
	// 	return
	// }

	fmt.Println("total", a+b)

}
