package main

import (
	"fmt"
	"math"
)

func Sin(x float64) float64 {
	return math.Sin(x)
}

func Cos(x float64) float64 {
	return math.Cos(x)
}

func main() {
	var a float64
	fmt.Scanln(&a)
	result1 := Sin(a)
	result2 := Cos(a)
	fmt.Println(result1)
	fmt.Println(result2)
}
