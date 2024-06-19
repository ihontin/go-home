package main

import (
	"fmt"
	"math"
)

// CompareRoundedValues возвращать булевое значение, которое будет истинным,
// если округленные значения равны, и разницу между округленными значениями.
func CompareRoundedValues(a, b float64, decimalPlaces int) (isEqual bool, difference float64) {
	decPlace := 1.0
	for i := 0; i < decimalPlaces; i++ {
		decPlace *= 10
	}

	x := math.Round(a*decPlace) / decPlace
	y := math.Round(b*decPlace) / decPlace
	fmt.Println(x, y)
	isEqual = x == y
	difference = math.Abs(x - y)
	return isEqual, difference
}

func main() {
	var a, b float64
	var dec int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&dec)
	fmt.Println(CompareRoundedValues(a, b, dec))
}
