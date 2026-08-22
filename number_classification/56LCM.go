package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	lcm := a
	if b > lcm {
		lcm = b
	}

	for {
		if lcm%a == 0 && lcm%b == 0 {
			break
		}
		lcm++
	}
	val := LCM(20, 30)
	fmt.Println("lcm i s", lcm, val)

}

// optimise approach
func GCD(a, b int) int {

	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	lcm := (a * b) / GCD(a, b)

	return lcm
}
