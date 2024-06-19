package main

import (
	"unicode/utf8"
)

//Напиши функцию countUniqueUTF8Chars(s string) int,
//которая будет подсчитывать количество уникальных UTF-8 символов в заданной строке.

func countUniqueUTF8Chars(s string) int {
	if utf8.RuneCountInString(s) < 1 {
		return 0
	}
	var count = make(map[int32]int)
	for _, b := range s {
		if _, ok := count[b]; !ok {
			count[b]++
		}
	}

	return len(count)
}

//func main() {
//	fmt.Println(countUniqueUTF8Chars("Hello, 世界!")) // Выводит false
//}
