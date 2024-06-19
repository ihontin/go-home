package main

import (
	"fmt"
)

// по убыванию
func sortDescInt(xs [8]int) [8]int {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == len(xs)-1 {
				continue
			}
			if xs[j] < xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}
	return xs
}

func sortAscInt(xs [8]int) [8]int {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == len(xs)-1 {
				continue
			}
			if xs[j] > xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}
	return xs
}

func sortDescFloat(xs [8]float64) [8]float64 {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == len(xs)-1 {
				continue
			}
			if xs[j] < xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}
	return xs
}

func sortAscFloat(xs [8]float64) [8]float64 {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == len(xs)-1 {
				continue
			}
			if xs[j] > xs[j+1] {
				xs[j], xs[j+1] = xs[j+1], xs[j]
			}
		}
	}
	return xs
}
func main() {
	xs := [8]int{1, 2, 3, 4, 5, 9, 8, 7}
	ysf := [8]float64{1.1, 2.2, 3.3, 4.4, 5.5, 9.9, 8.8, 7.7}

	fmt.Println("Sorted Int Array (Descending):", sortDescInt(xs))
	fmt.Println("Sorted Int Array (Ascending):", sortAscInt(xs))
	fmt.Println("Sorted Float Array (Descending):", sortDescFloat(ysf))
	fmt.Println("Sorted Float Array (Ascending):", sortAscFloat(ysf))
}
