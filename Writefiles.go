package main

import (
	"fmt"
	"os"
)

func writefiles(s string, content string) error {
	return os.WriteFile(s, []byte(content), 0644)
}

func main() {
	fmt.Println(writefiles("output.txt", "Go is powerful"))
}
