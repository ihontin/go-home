package main

import (
	"fmt"
)

func getSubSlice(xs []int, start, end int) []int {
	// Ваш код здесь
	return xs[start:end]
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subSlice := getSubSlice(numbers, 2, 6)
	fmt.Println(subSlice) // Вывод: [3 4 5 6]

}
