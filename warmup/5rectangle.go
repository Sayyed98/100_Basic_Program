package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var length, breadth float64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	length, _ = strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	breadth, _ = strconv.ParseFloat(scanner.Text(), 64)

	area := length * breadth
	fmt.Println(area)
	perimeter := 2 * (length + breadth)
	fmt.Println(perimeter)
}
