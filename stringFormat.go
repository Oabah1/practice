package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	var choice int

	fmt.Print("Enter word: ")
	fmt.Scan(&word)

	fmt.Println("1. Last char")
	fmt.Println("2. Capitalize")
	fmt.Println("3. Delete index")

	fmt.Print("Choose: ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		fmt.Println(string(word[len(word)-1]))

	case 2:
		fmt.Println(strings.ToUpper(word))

	case 3:
		var i int
		fmt.Print("Index: ")
		fmt.Scan(&i)

		if i >= 0 && i < len(word) {
			fmt.Println(word[:i] + word[i+1:])
		} else {
			fmt.Println("Invalid index")
		}
	}
}
