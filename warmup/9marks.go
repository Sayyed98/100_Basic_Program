package main

import "fmt"

func main() {

	fmt.Println("enter the Maths marks: ")
	var math int
	fmt.Scanln(&math)

	fmt.Println("enter the hindi marks: ")
	var hindi int
	fmt.Scanln(&hindi)

	fmt.Println("enter the english marks :")
	var english int
	fmt.Scanln(&english)

	fmt.Println("enter the physics marks: ")
	var physics int
	fmt.Scanln(&physics)

	fmt.Println("enter the chemistry marks: ")
	var chemistry int
	fmt.Scanln(&chemistry)

	total := math + hindi + english + physics + chemistry
	fmt.Println("total marks", total)

	average := (math + hindi + english + physics + chemistry) / 5

	fmt.Println("average marks: ", average)
}
