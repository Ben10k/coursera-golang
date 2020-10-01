package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("enter an integer to be added to the slice ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())
		if text == "X" {
			fmt.Println("input was [" + text + "], exiting")
			break
		}
		integer, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("unable to convert text [" + text + "] to integer, ignoring")
			continue
		}
		slice = append(slice, integer)
		sort.Ints(slice)
		fmt.Println(slice)
	}
}
