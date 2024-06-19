package main

import (
	"fmt"
	"strings"
)

// filterSentence , которая принимает предложение в виде строки и карту фильтра в формате map[string]bool.
// Функция должна удалить из предложения все слова, которые присутствуют в карте фильтра,
// и вернуть отфильтрованное предложение.
func filterSentence(sentence string, filter map[string]bool) string {
	var textList = strings.Fields(sentence)
	var formatedText = make([]string, 0, len(textList))
	for _, word := range textList {
		if !filter[word] {
			formatedText = append(formatedText, word)
		}
	}
	return strings.Join(formatedText, " ")
}

func main() {
	sentence := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	filter := map[string]bool{"ipsum": true, "elit": true}

	filteredSentence := filterSentence(sentence, filter)
	fmt.Println(filteredSentence)
}
