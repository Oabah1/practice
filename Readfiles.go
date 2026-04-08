package main

import (
	"fmt"
	"os"
)

func readfiles(s string) (string, error) {
	data, err := os.ReadFile(s)
	if err != nil {
		return " ", err
	}
	return string(data), err
}

func main() {
	fmt.Println(readfiles("input.txt"))
}
