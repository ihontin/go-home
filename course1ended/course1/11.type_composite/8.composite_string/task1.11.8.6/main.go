package main

import (
	"strings"
)

// CountVowels , которая будет принимать строку str и возвращать количество гласных букв в этой строке.
func CountVowels(s string) int {
	var vowels = []string{"а", "я", "у", "ю", "о", "е", "ё", "э", "и", "ы", "a", "e", "i", "o", "u", "y"}
	var count int
	for _, vow := range vowels {
		count += strings.Count(strings.ToLower(s), vow)
	}
	return count
}

//func main() {
//	// Пример использования функции CountVowels
//	count := CountVowels("Привет, мир!")
//	fmt.Println(count) // Вывод: 3
//}
