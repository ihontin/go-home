package main

import (
	"unicode/utf8"
)

// ReverseString (str string) string, которая принимает строку str и возвращает ее обратное представление.
func ReverseString(str string) string {
	if utf8.RuneCountInString(str) < 1 {
		return ""
	}
	return ReverseString(str[1:]) + str[:1]
}

//func main() {
//	fmt.Println(ReverseString("Hello, world!")) // Вывод: "!dlrow ,olleH"
//	fmt.Println(ReverseString("12345"))         // Вывод: "54321"
//}
