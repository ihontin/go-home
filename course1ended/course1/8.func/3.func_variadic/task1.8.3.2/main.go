package main

import (
	"fmt"
	"strings"
)

//Функция должна объединять строки через разделитель в порядке их следования и возвращать строку в формате
//“even: some string, odd: some string”,где even - количество строк с четным индексом,
//odd - количество строк с нечетным индексом. Индексация начинается с 0.

func ConcatenateStrings(sep string, allStrings ...string) string {
	var strFul, strEven, strOdd string
	var even, odd int
	for i, someStr := range allStrings {
		if (i+2)%2 == 0 {
			even++
			strEven += fmt.Sprintf("%s%s", someStr, sep)
		} else {
			odd++
			strOdd += fmt.Sprintf("%s%s", someStr, sep)
		}
	}
	strEven = strings.TrimSuffix(strEven, sep)
	strOdd = strings.TrimSuffix(strOdd, sep)
	strFul = "even: " + strEven + ", " + "odd: " + strOdd
	return strFul
}
func main() {
	fmt.Println(ConcatenateStrings("*-*", "q", "E", "q", "E", "q", "E", "q", "E", "q", "E", "q",
		"E", "q", "E", "q", "E", "q", "E", "q", "E", "q", "E", "q"))
}
