package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println("arr[2:6] =",arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])

	fmt.Println("modify...")

	s1 := arr[2:6]
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("Slice extend...")
	arr = [...]int{0, 1, 2, 3, 4, 5, 6}
	a1 := arr[2:6]
	a2 := a1[3:5]
	fmt.Println("a1=", a1)
	fmt.Println("a2=", a2)

	// index 下标不能超过len, re-slice 不能超过cap
	fmt.Println(len(a1), cap(a1))
	fmt.Println(len(a2), cap(a2))
}
