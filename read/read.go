package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	file, _ := os.Open(readFileName())
	slice := readFile(file)
	_ = file.Close()

	printResults(slice)
}

func readFileName() string {
	fmt.Println("Enter the input file name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readFile(file *os.File) []name {
	slice := make([]name, 0, 3)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		slice = append(slice, line2Name(fileScanner.Text()))
	}
	return slice
}

func line2Name(text string) name {
	line := strings.Split(strings.TrimSpace(text), " ")
	return name{fname: line[0], lname: line[1]}
}

func printResults(slice []name) {
	fmt.Println("Results:")
	for _, n := range slice {
		fmt.Printf("First name [%s] Last name [%s]\n", n.fname, n.lname)
	}
}
