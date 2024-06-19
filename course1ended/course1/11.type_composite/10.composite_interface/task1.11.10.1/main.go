package main

import (
	"fmt"
)

// getType(i interface{}) string, которая принимает на вход объект типа interface{} и возвращает его тип в виде строки.
func getType(i interface{}) string {
	switch v := i.(type) {
	case string:
		return fmt.Sprintf("%T", v)
	case int:
		return fmt.Sprintf("%T", v)
	case []int:
		return fmt.Sprintf("%T", v)
	default:
		return "Пустой интерфейс"
	}
}

//func main() {
//	var i interface{} = 42
//	fmt.Println(getType(i)) // Вывод: "int"
//
//	var j interface{} = "Hello, World!"
//	fmt.Println(getType(j)) // Вывод: "string"
//
//	var k interface{} = []int{1, 2, 3}
//	fmt.Println(getType(k)) // Вывод: "[]int"
//
//	var l interface{} = interface{}(nil)
//	fmt.Println(getType(l)) // Вывод: "Пустой интерфейс"
//}
