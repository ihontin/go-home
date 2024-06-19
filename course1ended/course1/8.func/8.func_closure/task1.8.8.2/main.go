package main

import "fmt"

// multiplier, которая будет принимать в качестве аргумента factor типа float64 и возвращать функцию,
// которая будет умножать переданное число на factor.
func multiplier(factor float64) func(float64) float64 {
	return func(f float64) float64 {
		return factor * f
	}
}

func main() {
	// Пример использования функции multiplier
	m := multiplier(2.5)
	result := m(10)
	result2 := m(20)
	result3 := m(5)
	fmt.Println(result, result2, result3) // Вывод: 25
	// ЗАМЫКАНИЕ
	//step1 := func(x int) func(int) func(int) int {
	//	return func(n int) func(int) int {
	//		return func(t int) int {
	//			return n * x * t
	//		}
	//	}
	//}
	//var (
	//	a = 3
	//	b = 5
	//	c = 2
	//)
	//fmt.Println(step1(a)(b)(c))
	//
	//var step2 = step1(a)
	//fmt.Println(step2(b)(c))
	//
	//var step3 = step2(b)
	//fmt.Println(step3(c))
}
