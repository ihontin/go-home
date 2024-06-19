package main

import "fmt"

//func main() {
//	var num int = 10
//	var str string = "Hello"
//
//	fmt.Println(getVariableType(num)) // Вывод: "int"
//	fmt.Println(getVariableType(str)) // Вывод: "string"
//}

// которая будет возвращать тип переменной с использованием функции fmt.Sprintf().
// Функция должна принимать переменную в качестве аргумента и возвращать ее тип в виде строки.
func getVariableType(variable interface{}) string {
	return fmt.Sprintf("%T", variable)

}
