package main

import (
	"fmt"
)

func count2(s string) int {
	return len(s)
}

func main() {
	fmt.Println(count2("go go"))
	fmt.Println(count2("exercise exercise"))
	fmt.Println(count2("lang"))

}
