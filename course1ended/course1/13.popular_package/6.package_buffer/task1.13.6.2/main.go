package main

import (
	"bufio"
	"bytes"
)

//Необходимо написать функцию getScanner, которая принимает в качестве аргумента объект типа bytes.Buffer,
//содержащий строковые данные, и возвращает объект типа bufio.Scanner.

func getScanner(b *bytes.Buffer) *bufio.Scanner {
	return bufio.NewScanner(b)
}

//func main() {
//	// Create a buffer with some data
//	data := []byte("Hello\n,\n World!")
//	buffer := bytes.NewBuffer(data)
//
//	// Call the getScanner function
//	scanner := getScanner(buffer) // сканер позволяет читать текст по частям, в цикле по строкам, словам, байтам, рунам
//
//	// Verify that the returned reader is not nil
//	if scanner == nil {
//		panic("Expected non-nil reader, got nil")
//	}
//	for scanner.Scan() {
//		println(scanner.Text())
//	}
//}

//Hello
//,
//World!
