package main

import (
	"fmt"
	"math"
)

func hypotenuse(a, b float64) float64 {
	// реализация функции
	return math.Sqrt(a*a + b*b)
}

func main() {
	var x, y float64
	fmt.Scanln(&x, &y)
	fmt.Println(hypotenuse(x, y))
}
