package main

import "fmt"

// adder должна принимать один аргумент initial, который будет являться начальным значением для сложения.
// Функция adder должна возвращать функцию, которая будет принимать один аргумент value и возвращать сумму initial и value.
func adder(initial int) func(int) int {
	return func(i int) int {
		return initial + i
	}
	// ваш код здесь
}

func main() {

	// пример использования функции adder
	addTwo := adder(2)
	result := addTwo(3)
	result2 := addTwo(4)
	result3 := addTwo(5)
	fmt.Println(result, result2, result3) // выводит 5 6 7

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
