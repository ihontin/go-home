package main

import (
	"fmt"
	"math"
)

func Floor(x float64) float64 {
	return math.Floor(x)
}

func main() {
	var a float64
	fmt.Scanln(&a)
	result := Floor(a)
	fmt.Println(result)
}
