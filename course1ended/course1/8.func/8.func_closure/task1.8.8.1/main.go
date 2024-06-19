package main

import "fmt"

// createCounter
// которая будет возвращать функцию, которая в свою очередь будет возвращать следующее число при каждом вызове.
// Для реализации данной функции необходимо использовать замыкание.
func createCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {
	counter := createCounter()
	fmt.Println(counter()) // 1
	fmt.Println(counter()) // 2
	fmt.Println(counter()) // 3
	counter2 := createCounter()
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())

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
