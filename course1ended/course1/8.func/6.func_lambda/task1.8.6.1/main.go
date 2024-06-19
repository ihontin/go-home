package main

import "fmt"

func Sum(a ...int) int {
	var sumA int
	for i, val := range a {
		if i == 0 {
			sumA = val
			continue
		}
		sumA += val
	}
	return sumA
}

func Mul(a ...int) int {
	var sumA int
	for i, val := range a {
		if i == 0 {
			sumA = val
			continue
		}
		sumA *= val
	}
	return sumA
}

func Sub(a ...int) int {
	var subA int
	for i, val := range a {
		if i == 0 {
			subA = val
			continue
		}
		subA -= val
	}
	return subA
}

func MathOperate(op func(a ...int) int, a ...int) int {
	return op(a...)
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))  // Output: 5
	fmt.Println(MathOperate(Mul, 1, 7, 3))  // Output: 21
	fmt.Println(MathOperate(Sub, 13, 2, 3)) // Output: 8
}
