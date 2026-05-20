package main

import (
	"fmt";
	"learn-go/mydict";
)

func main() {
	dictionary := mydict.Dictionary{}

	// =======================================

	// def, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(def)
	// }

	// =======================================

	// word := "hello"
	// definition := "Greeting"

	// err := dictionary.Add(word, definition)
	// if (err != nil) {
	// 	fmt.Println(err)
	// }

	// err2 := dictionary.Add(word, definition)
	// if (err2 != nil) {
	// 	fmt.Println(err2)
	// }

	word := "hello"
	dictionary.Add(word, "first")
	err := dictionary.Update("word", "second")
	if err != nil {
		fmt.Println(err)
	}
	dictionary.Delete(word)
	findword, err2 := dictionary.Search(word)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(findword)

}
