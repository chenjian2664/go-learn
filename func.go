package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int  {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func swap(a, b  *int) {
	*a, *b = *b, *a
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d %d) \nresult is :", opName, a, b)
	return op(a, b)
}

func main() {
	fmt.Println(eval(3, 4, "*"))
	//fmt.Println(eval(3, 4, "aa"))
	q, r := div(19, 3)
	fmt.Println(q, r)

	fmt.Println(apply(func(a, b int ) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	swap(&q, &r)
	fmt.Println(q, r)
}
