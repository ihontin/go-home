package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	var a float64
	fmt.Scanln(&a)
	result := Sqrt(a)
	fmt.Println(result)
}
