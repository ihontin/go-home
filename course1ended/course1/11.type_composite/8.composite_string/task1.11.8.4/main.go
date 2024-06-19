package main

// concatStrings, которая будет принимать произвольное количество строк и возвращать их конкатенацию в виде одной строки.
func concatStrings(xs ...string) string {
	var newStr string
	for _, word := range xs {
		newStr += word
	}
	return newStr
}

//func main() {
//	result := concatStrings("Hello", " ", "world!")
//	fmt.Println(result) // Вывод: "Hello world!"
//}
