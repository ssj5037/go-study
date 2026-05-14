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

func canIDrink(age int) bool {
	if koreanAge:= age + 2; koreanAge < 18 {
		return false
	}
	return true
}

func canIDrink2(age int) bool {
	switch koreanAge:= age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	default:
		return false
	}
}

type person struct {
	name string
	age int
	favFood []string
}

func mainTheory() {
	fmt.Println(multiply(2,2))
	totalLength, upperName := lenAndUpper("sujin")
	fmt.Println(totalLength, upperName)
	repeatMe("sujin", "kim", "park", "lee")
	total := superAdd(1,2,3,4,5,6,7,8,9,10)
	fmt.Println(total)

	fmt.Println(canIDrink(15))
	fmt.Println(canIDrink(20))
	fmt.Println(canIDrink2(16))
	fmt.Println(canIDrink2(20))

	a := 2
	b := a
	a = 10
	fmt.Println(a, b) // 10 2

	c := 2
	d := 5
	fmt.Println(&c, &d) // 0x49fdb8baa110 0x49fdb8baa118

	e := 2
	f := &e
	fmt.Println(e, f) // 2 0x5f111b588050
	fmt.Println(&e, f) // 0x5f111b588050 0x5f111b588058
	fmt.Println(*f) // 2
	*f = 10
	fmt.Println(e, *f) // 10 10

	names := []string{"nico", "lynn", "dal", "mark", "flynn"}
	names = append(names, "jason")
	fmt.Println(names)

	nico := map[string]string{"name": "nico", "age": "18"}
	for key, value := range nico {
		fmt.Println(key, value)
	}

	favFood := []string{"kimchi", "ramen"}
	nicoPerson := person{name:"nico", age:18, favFood:favFood}
	fmt.Println(nicoPerson.name)
	
}

