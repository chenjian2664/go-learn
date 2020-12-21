package main

import "fmt"

func main() {
	m := map[string]string {
		"hello": "world",
		"ni": "hao",
		"learn": "go",
	}

	fmt.Println(m)

	var nMap map[string]int
	nMap = map[string]int{}
	fmt.Println(nMap)

	var m2 map[int]int
	m2 = map[int]int{}
	fmt.Println(m2)

	v1 := m["hello"]
	v2 := m["ssss"]
	fmt.Println(v1, v2)

	fmt.Println("exists key ?")
	v3, ok := m["hello"]
	fmt.Println(v3, ok)
	v4, ok := m["xxxx"]
	fmt.Println(v4, ok)

	m["aaa"] = "ddd"

	for k, v := range m {
		fmt.Println(k, v)
	}
}
