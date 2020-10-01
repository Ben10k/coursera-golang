package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var person = make(map[string]string)

	fmt.Println("enter a name:")
	person["name"] = readLine(scanner)

	fmt.Println("enter an address:")
	person["address"] = readLine(scanner)

	result, _ := json.Marshal(person)

	fmt.Println(string(result))
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}
