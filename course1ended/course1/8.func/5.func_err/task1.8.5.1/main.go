package main

import (
	"fmt"
	"strconv"
)

// CalculatePercentageChange (initialValue, finalValue string) (float64, error),
// которая вычисляет процентное изменение между двумя значениями и возвращает процентное изменение и ошибку.
// — Функция должна возвращать ошибку, если initialValue или finalValue не являются числом.
// — Функция должна вернуть 0 в случае деления на 0.
func CalculatePercentageChange(initialValue, finalValue string) (float64, error) {
	firstVal, err := strconv.ParseFloat(initialValue, 64)
	if err != nil {
		return 0, err
	}
	if firstVal == 0 {
		return 0, nil
	}
	lastVal, err2 := strconv.ParseFloat(finalValue, 64)
	if err2 != nil {
		return 0, err2
	}
	return ((lastVal - firstVal) / firstVal) * 100, nil
}

func main() {
	fmt.Println(CalculatePercentageChange("0", "212e"))
}
