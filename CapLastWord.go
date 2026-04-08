package main

import (
	"fmt"
	"strings"
)

func CapLastWord(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" && i > 0 {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i--
		}
	}
	return s
}

func main() {
	fmt.Println(CapLastWord([]string{"God", "is", "good", "(up)"}))
}
