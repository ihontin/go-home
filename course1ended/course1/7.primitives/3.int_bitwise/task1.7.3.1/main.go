package main

import "fmt"

// — Функция bitwiseAnd должна возвращать результат операции и
func bitwiseAnd(a, b int) int {
	return a & b
}

// —Функция bitwiseOr должна возвращать результат операции или
func bitwiseOr(a, b int) int {
	return a | b
}

// — Функция bitwiseXor должна возвращать результат операции xor
func bitwiseXor(a, b int) int {
	return a ^ b
}

// — Функция bitwiseLeftShift должна возвращать результат операции левый битовый сдвиг
func bitwiseLeftShift(a, b int) int {
	return a << b
}

// — Функция bitwiseRightShift должна возвращать результат операции правый битовый сдвиг
func bitwiseRightShift(a, b int) int {
	return a >> b
}

func main() {
	var a int
	var b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println("a & b =", bitwiseAnd(a, b))
	fmt.Println("a | b =", bitwiseOr(a, b))
	fmt.Println("a ^ b =", bitwiseXor(a, b))
	fmt.Println("a << b =", bitwiseLeftShift(a, b))
	fmt.Println("a >> b =", bitwiseRightShift(a, b))

}
