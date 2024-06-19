package main

import "fmt"

// CalculatePercentageChange (initialValue, finalValue float64) float64,
// которая принимает начальное и конечное значения и вычисляет процентное изменение между ними.
func CalculatePercentageChange(initialValue, finalValue float64) float64 {
	return ((finalValue - initialValue) / initialValue) * 100
}

func main() {
	fmt.Println(CalculatePercentageChange(111, 112))
}
