package main

import "fmt"

// package variables
var (
	aa = 3
	bb = 4
	ss = "aaa"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "aaa"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "ddd"
	fmt.Println(a, b, c, s)
}

func variableShoter() {
	a, b, c, s := 3, 4, true, "aaa"
	fmt.Println(a, b, c, s)
}

func consts() {
	const f string = "aa"
	const a, b  = 1, 23
	fmt.Println(f, a, b)
}

func enums() {
	const (
		a = iota
		aa
		c
		d
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(a, aa, c, d)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShoter()
	fmt.Println(aa, bb, ss)

	consts()
	enums()
}
