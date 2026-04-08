package main

import (
	"fmt"
	"strings"
)

func count(s string) map[string]int {
	word := strings.Fields(s)
	count := make(map[string]int)
	for _, i := range word {
		count[i]++
	}
	return count
}

func main() {
	for k, v := range count("go go lang exercise exercise") {
		fmt.Printf("%s: %d\n", k, v)
	}
}
