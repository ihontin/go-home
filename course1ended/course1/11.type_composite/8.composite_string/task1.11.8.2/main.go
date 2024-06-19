package main

import (
	"reflect"
	"unsafe"
)

// Напиши функцию getStringHeader(s string) reflect.StringHeader,
// которая будет возвращать заголовок среза для строки s с использованием unsafe.Pointer.
func getStringHeader(s string) reflect.StringHeader {
	newS := (*reflect.StringHeader)(unsafe.Pointer(&s))
	return *newS
}

//func main() {
//	s := "Hello, World!"
//	header := getStringHeader(s)
//	fmt.Printf("Data: %v\n", header.Data)
//	fmt.Printf("Len: %v\n", header.Len)
//	fmt.Println(len(s))
//}
