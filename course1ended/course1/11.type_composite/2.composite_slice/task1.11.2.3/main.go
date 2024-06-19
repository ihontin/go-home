package main

import "fmt"

func bitwiseXOR(n, res int) int {
	result := n ^ res
	if result == 0 {
		return 0
	}
	return result
}

func findSingleNumber(numbers []int) int {
	// ваш код

	for i, x := range numbers {
		for j, y := range numbers {
			if i == j {
				continue
			}
			v := bitwiseXOR(y, x)
			if v == 0 {
				break
			} else if v != 0 && j != len(numbers)-1 {
				continue
			} else if v != 0 && j == len(numbers)-1 {
				return x
			}
		}
	}
	return 0
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	singleNumber := findSingleNumber(numbers)
	fmt.Println(singleNumber) // 5
}
