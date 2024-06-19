package main

import "fmt"

// DivideAndRemainder (a, b int) (int, int), которая принимает два целых числа
// и возвращает частное и остаток от деления первого числа на второе.
func DivideAndRemainder(a, b int) (int, int) {
	var res1, res2 int
	if b == 0 {
		res1 = 0
		res2 = 0
	} else {
		res1 = a / b
		res2 = a % b
	}
	return res1, res2
}

func main() {
	fmt.Println(DivideAndRemainder(122, 8))
}
