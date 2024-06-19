package main

import (
	"strings"
)

// Напиши функцию, которая создает текст с уникальными словами, используя карту (map).
// Функция должна принимать текст в виде строки и возвращать новую строку,
// содержащую только уникальные слова из исходного текста, сохраняя порядок слов.
func getUniqueWords(text string) string {
	var listText = strings.Fields(text)
	var uniqListText = make([]string, 0, len(listText))
	var mapUniqText = make(map[string]bool)
	for _, word := range listText {
		if _, ok := mapUniqText[word]; !ok {
			mapUniqText[word] = true
			uniqListText = append(uniqListText, word)
		}
	}
	return strings.Join(uniqListText, " ")
}

//func main() {
//	newStr := createUniqueText("bar bar bar foo foo baz")
//	fmt.Println(newStr)
//}
