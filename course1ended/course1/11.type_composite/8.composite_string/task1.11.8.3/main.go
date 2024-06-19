package main

//Необходимо написать две функции: getBytes(s string) []byte и getRunes(s string) []rune.
//	Обе функции должны принимать строку s в качестве аргумента и возвращать соответственно срез байтов и срез символов.

// Функция getBytes возвращает срез байтов строки s
func getBytes(s string) []byte {
	return []byte(s)
}

// Функция getRunes возвращает срез символов строки s
func getRunes(s string) []rune {
	return []rune(s)
}

//func main() {
//	var a = "fffff"
//	var x = getRunes(a)
//	var y = getBytes(a)
//	fmt.Println(x, y)
//}
