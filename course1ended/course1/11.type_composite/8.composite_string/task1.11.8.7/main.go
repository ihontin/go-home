package main

import (
	"fmt"
	"strings"
)

// ReplaceSymbols , которая заменяет все вхождения символа old в строке str на символ new и возвращает полученную строку.
// Функция должна быть реализована на языке программирования Golang.
func ReplaceSymbols(s string, old, new rune) string {
	var newS = strings.Replace(s, string(old), string(new), -1)
	//newS = string(markdown.ToHTML([]byte(newS), nil, nil))
	return newS
}

func main() {
	result := ReplaceSymbols("Hello, world!", 'o', '0')
	fmt.Println(result) // result должно быть равно "Hell0, w0rld!"
}
