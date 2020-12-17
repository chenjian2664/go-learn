package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	r := ""
	switch  {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		r = "F"
	case score < 80:
		r = "C"
	case score < 90:
		r = "B"
	case score <= 100:
		r = "A"
	}
	return r
}

func main() {
	const filename  = "abc"
	if constents, err := ioutil.ReadFile("helloworld.go"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(constents))
	}

	fmt.Println(
		grade(0),
		grade(10),
		grade(65),
		grade(87),
		//grade(111),
		)
}
