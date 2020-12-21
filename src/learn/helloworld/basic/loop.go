package basic

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func converToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("aaaa")
	}
}

func main() {
	fmt.Println(
		converToBin(5),
		converToBin(13),
		converToBin(123123123))

	printFile("helloworld.go")
	forever()
}
