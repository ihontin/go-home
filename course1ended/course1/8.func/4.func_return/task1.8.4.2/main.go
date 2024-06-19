package main

import "fmt"

// FindMaxAndMin которая принимает неопределенное количество целочисленных аргументов
// и возвращает два значения: максимальное и минимальное значение из переданных аргументов.
func FindMaxAndMin(n ...int) (int, int) {
	var (
		max int
		min int
	)
	for i, num := range n {
		if i == 0 {
			max = num
			min = num
		}
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return max, min
}
func main() {
	//Функция должна находиться в пакете main.

	fmt.Println(FindMaxAndMin(1, 5, 4, 6, 7, -9, 8, 99, 81, 9))
}
