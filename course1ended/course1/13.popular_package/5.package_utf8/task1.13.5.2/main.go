package main

import (
	"unicode"
)

//Напиши функцию countRussianLetters(s string) map[rune]int
//которая принимает строку s и возвращает карту, содержащую количество каждой русской буквы в строке.

func isRussianLetter(r rune) bool {
	if r == 1025 || r == 1105 || r > 1039 && r < 1104 {
		return true
	}
	return false
}

func countRussianLetters(s string) map[rune]int {

	counts := make(map[rune]int)
	if s == "" {
		return counts
	}
	for _, char := range s {
		if isRussianLetter(char) {
			counts[unicode.ToLower(char)]++
		}
	}
	return counts
}

//func main() {
//	result := countRussianLetters("Привет, мир!")
//	for key, value := range result {
//		fmt.Printf("%c: %d ", key, value) // в: 1 е: 1 т: 1 м: 1 п: 1 р: 2 и: 2
//	}
//}
