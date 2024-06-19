package main

import (
	"fmt"
	"strings"
)

// countWordOccurrences, которая принимает текст в качестве аргумента и возвращает карту (map)
//с подсчитанным количеством вхождений каждого слова в тексте.

func countWordOccurrences(text string) map[string]int {
	var listText = strings.Fields(text)
	var occurre = make(map[string]int)
	for _, key := range listText {
		occurre[key]++
	}
	return occurre
}

func main() {
	text := "Lorem ipsum dolor sit amet consectetur adipiscing elit ipsum"
	occurrences := countWordOccurrences(text)

	for word, count := range occurrences {
		fmt.Printf("%s: %d\n", word, count)
	}
}
