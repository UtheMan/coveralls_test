package main

import "fmt"

func addNumbers(a, b int) int {
	return a + b
}

func subtractNumbers(a, b int) int {
	return a - b
}

func main() {
	s := addNumbers(1, 2)
	d := subtractNumbers(1, 2)
	fmt.Printf("sum is %d, difference is %d", s, d)
}
