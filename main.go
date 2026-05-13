package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	fmt.Println(multiply(2,2))
	totalLength, upperName := lenAndUpper("sujin")
	fmt.Println(totalLength, upperName)
	repeatMe("sujin", "kim", "park", "lee")
	total := superAdd(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(total)
}