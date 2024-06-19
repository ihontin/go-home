package main

import (
	"strings"
)

//Напиши функцию CountWordsInText(txt string, words []string) map[string]int,
//которая будет подсчитывать количество вхождений каждого слова из заданного списка в тексте.

//При подсчете слов в тексте необходимо игнорировать регистр с помощью функции strings.ToLower().
//Для разделения текста на слова используй функцию strings.Fields().

func CountWordsInText(txt string, words []string) map[string]int {
	var uniqWords = make(map[string]int)
	if txt == "" || len(words) == 0 {
		return uniqWords
	}
	replacer := strings.NewReplacer(".", "", ",", "")
	var WordsList = strings.Fields(strings.ToLower(replacer.Replace(txt)))

	for _, ws := range words {
		uniqWords[ws] = 0
	}
	for _, w := range WordsList {
		if _, ok := uniqWords[w]; ok {
			uniqWords[w]++
		}
	}
	return uniqWords
}

//func main() {
//	txt := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a diam lectus. Sed sit amet ipsum mauris.
//        Maecenas congue ligula ac quam viverra nec consectetur ante hendrerit. Donec et mollis dolor.
//        Praesent et diam eget libero egestas mattis sit amet vitae augue.`
//	words := []string{"sit", "amet", "lorem"}
//
//	result := CountWordsInText(txt, words)
//
//	fmt.Println(result) // map[amet:2 lorem:1 sit:3]
//}
