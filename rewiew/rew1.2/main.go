package main

import (
	"fmt"
	"math"
)

// 1
// Написать программу fizz buzz

// Если число делится на 3, то вывести fizz
// Если число делится на 5, то вывести buzz
// Если число делится на 3 и на 5, то вывести fizz buzz

// Покрыть табличными тестами
func fizzBuzz(x int) string {
	switch {
	case x%3 == 0 && x%5 == 0:
		return "fizz buzz"
	case x%3 != 0 && x%5 == 0:
		return "buzz"
	case x%3 == 0 && x%5 != 0:
		return "fizz"
	default:
		return "wrong number"
	}
}

// Задача 2
// Написать бенчмарк для Fibonacci, двух реализаций, при помощи рекурсии и формулы Бине

func FibbRecurs(sliceFib int) int {
	if sliceFib < 2 {
		return sliceFib
	}
	return FibbRecurs(sliceFib-1) + FibbRecurs(sliceFib-2)
}
func FibonacciFormula(n int) int {
	phi := (1 + math.Sqrt(5)) / 2
	return int(math.Round(math.Pow(phi, float64(n)) / math.Sqrt(5)))
}

// Задача 3
// Написать функцию округления float, используя math.Round и math.Pow
func roundNumber(x float64, round int) float64 {
	fa := math.Pow(10, float64(round))
	answer := math.Round(x*fa) / fa
	return answer
}

func main() {
	fbNum := 35
	fmt.Println(fizzBuzz(fbNum))

	numFibonacci := 6
	fmt.Println(FibonacciFormula(numFibonacci), FibbRecurs(numFibonacci))

	num := 3251.1734
	round := 1
	fmt.Println(roundNumber(num, round))

}
