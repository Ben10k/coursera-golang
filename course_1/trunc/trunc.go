package main

import (
	"fmt"
)

func main() {
	var input float64

	fmt.Println("enter a floating point number:")
	_, _ = fmt.Scanf("%f", &input)

	fmt.Printf("trumcated number: %d\n", int(input))
}
