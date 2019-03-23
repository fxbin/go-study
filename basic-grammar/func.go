package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
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
		panic("unsupported operation:" + op)
	}
	return result
}

// 13 / 3 = 4 ... 1
//func div(a, b int) (int, int)  {
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func apply(op func(int, int) int, a, b int) int {
	fmt.Printf("Calling funcation %s with args (%d, %d)",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
		a, b)
	return op(a, b)
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swap1(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println(
		eval(1, 2, "+"),
		eval(1, 2, "-"),
		eval(1, 2, "*"),
		eval(1, 2, "/"),
	)

	q, r := div(13, 3)
	fmt.Print(q, r)
	fmt.Println(div(13, 3))

	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(
			float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	a, b = swap1(a, b)
	fmt.Println(a, b)
}
