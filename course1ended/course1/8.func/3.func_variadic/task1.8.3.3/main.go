package main

import "fmt"

// PrintNumbers будет принимать произвольное количество целых чисел
// и выводить их на экран по одному в строке с помощью функции fmt.Println.
func PrintNumbers(randnums ...int) {
	for _, iterNum := range randnums {
		fmt.Println(iterNum)
	}
}

func main() {
	PrintNumbers(1, 2, 3, 4, 5)

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
