package main

import "fmt"

func calculation(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func main() {
	fmt.Println(
		calculation(1, 2, "+"),
		calculation(1, 2, "-"),
		calculation(1, 2, "*"),
		calculation(1, 2, "/"),
	)
}
