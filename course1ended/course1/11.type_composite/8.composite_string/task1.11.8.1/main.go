package main

import (
	"unicode/utf8"
)

// countBytes (s string) int должна принимать строку s и возвращать количество байтов в этой строке.
func countBytes(s string) int {
	return len(s)
}

// countSymbols (s string) int должна принимать строку s и использовать функцию
// utf8.RuneCountInString для подсчета количества символов в этой строке.
func countSymbols(s string) int {
	return utf8.RuneCountInString(s)
}

//func main() {
// Пример использования функции countBytes
//bytes := countBytes("Привет, мир!")
//bytes := countBytes("Привет, мир!")

// Пример использования функции countSymbols
//symbols := countSymbols("Привет, мир!")
//symbols := countSymbols("Привет, мир!")
//}
