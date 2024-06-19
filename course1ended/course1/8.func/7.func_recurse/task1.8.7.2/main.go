package main

import "fmt"

// Fibonacci (n int) int, которая вычисляет n-ый элемент ряда Фибоначчи рекурсивно.
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println(Fibonacci(4))
}
