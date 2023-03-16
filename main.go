package main

import "fmt"

func addNumbers(a, b int) int {
	return a + b
}

func subtractNumbers(a, b int) int {
	return a - b
}

func multiplyNumbers(a, b int) int {
	return a * b
}

func main() {
	s := addNumbers(1, 2)
	d := subtractNumbers(1, 2)
	m := multiplyNumbers(1, 2)
	fmt.Printf("sum is %d, difference is %d, multiplied %d", s, d, m)
	fmt.Println("test PR")
}
