package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",s, len(s), cap(s))
}

func main() {
	var s []int
	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s, i)
	}
	fmt.Println(s)

	fmt.Println("make slice")
	s2 := make([]int, 16)
	s3 := make([]int, 10, 23)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("copy slice")
	printSlice(s)
	copy(s2, s)
	printSlice(s2)

	fmt.Println("append slice")
	s4 := append(s2[:3], s2[4:]...)
	printSlice(s4)

	fmt.Println("remove head")
	fmt.Println("before remove head", s4)
	s4 = s4[1:]
	fmt.Println("after remove head", s4)

	fmt.Println("remove tail")
	fmt.Println("before remove tail", s4)
	s4 = s4[:len(s4) - 1]
	fmt.Println("after remove tail", s4)
}
