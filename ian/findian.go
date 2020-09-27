package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("enter a string:")

	lowerCaseInput := scanLineToLowerCase()

	if isIan(lowerCaseInput) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}

func scanLineToLowerCase() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.ToLower(strings.TrimSpace(scanner.Text()))
}

func isIan(lowerInput string) bool {
	return strings.HasPrefix(lowerInput, "i") &&
		strings.HasSuffix(lowerInput, "n") &&
		strings.Contains(lowerInput, "a")
}
