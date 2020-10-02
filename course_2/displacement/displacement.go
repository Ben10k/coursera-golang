package main

import (
	"fmt"
	"math"
)

func main() {
	a := readFloat("a")
	v0 := readFloat("v0")
	s0 := readFloat("s0")

	fn := GenDisplaceFn(a, v0, s0)

	t := readFloat("t")

	fmt.Printf("after t[%v], the dispalacement is [%v]\n", t, fn(t))
}

func readFloat(name string, ) float64 {
	var temp float64
	fmt.Printf("enter %v:\n", name)
	_, _ = fmt.Scanf("%f", &temp)
	return temp
}

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return ((a * math.Pow(t, 2)) / 2) + (v0 * t) + s0
	}
}
