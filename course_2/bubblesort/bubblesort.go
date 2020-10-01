package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 10)
	slice = readIntegers(slice)

	BubbleSort(slice)
	log.Println("sorted slice:", slice)
}

func readIntegers(slice []int) []int {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 10; i++ {
		log.Println("enter an integer to be added to the slice, enter X to stop adding to the slice ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())
		if text == "X" {
			break
		}
		integer, err := strconv.Atoi(text)
		if err != nil {
			log.Println("unable to convert text [" + text + "] to integer, ignoring")
			continue
		}
		slice = append(slice, integer)
	}
	log.Println("no more integers can be added, the current slice is", slice)

	return slice
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < (len(slice) - i - 1); j++ {
			if (slice)[j] > (slice)[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, i int) {
	(slice)[i], (slice)[i+1] = (slice)[i+1], (slice)[i]
}
