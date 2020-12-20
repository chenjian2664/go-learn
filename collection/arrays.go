package main

import "fmt"

func printArray(arr [3]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [10]int
	arr2 := [3]int {1, 23, 4}
	arr3 := [...]int{1,2,3,3,4,1,1}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int
	fmt.Println(grid)

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for _, __ := range arr3 {
		fmt.Println(__)
	}

	printArray(arr2)
	// array 是值类型， [3]int 和 [5]int是两个不同的类型
	//printArray(arr3)
}
