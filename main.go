package main

import (
	"fmt"

	"github.com/sehoon123/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Print(err)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "definition :", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Print(err2)
	}
}
