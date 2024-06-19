package main

import (
	"fmt"
	"math/big"
	"strings"
)

// Factorial должна вычислять факториал числа и возвращать результат.
func Factorial(n *int) int {
	result := big.NewInt(1)

	for i := 1; i <= *n; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}

	return int(result.Int64())
}

// isPalindrome должна проверить, является ли строка палиндромом, и вернуть результат проверки.
func isPalindrome(str *string) bool {
	var (
		origStr   = strings.ToLower(*str)
		reversStr string
	)
	for _, letter := range origStr {
		reversStr = string(letter) + reversStr
	}
	return origStr == reversStr
}

// CountOccurrences должна вернуть количество вхождений числа в срез.
func CountOccurrences(numbers *[]int, target *int) int {
	var count int
	for _, n := range *numbers {
		if n == *target {
			count++
		}
	}
	return count
}

// ReverseString должна развернуть строку и вернуть результат.
func ReverseString(str *string) string {
	reversStr := ""
	for _, letter := range *str {
		reversStr = string(letter) + reversStr
	}
	return reversStr
}

func main() {
	var a = 23
	fmt.Println("fac off 23 -", Factorial(&a))

	var (
		s1 = "kragomis1)6"
		s2 = "123&palindromordnilap&321"
	)
	fmt.Println("is it false -", isPalindrome(&s1), "\nis it true -", isPalindrome(&s2))

	countIntrence := []int{5, 5, 5, 5, 23, 9, 523, 23, 23, 2359523}
	fmt.Println("it should be 3 -", CountOccurrences(&countIntrence, &a))

	fmt.Println("revers of", s1, "is -", ReverseString(&s1))
}
