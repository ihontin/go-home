package main

import "fmt"

// Сложение - Вычитание - Умножение - Деление - Остаток от деления
func calculate(a int, b int) (int, int, int, int, int) {
	c := int(a / b)
	return a + b, a - b, a * b, c, a % b
}

func main() {

	var a int
	var b int
	var sum int
	var difference int
	var product int
	var quotient int
	var remainder int

	a = 10
	b = 3
	sum = 13
	difference = 7
	product = 30
	quotient = 3
	remainder = 1
	sum, difference, product, quotient, remainder = calculate(a, b)
	fmt.Printf("a = %d b = %d sum = %d difference = %d product = %d quotient = %d remainder = %d", a, b, sum, difference, product, quotient, remainder)
}
