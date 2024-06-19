package main

import (
	"bufio"
	"bytes"
	"fmt"
)

//Необходимо написать функцию getReader(b *bytes.Buffer) *bufio.Reader,
//которая принимает объект bytes.Buffer с данными в виде строки и возвращает объект bufio.Reader.

func getReader(b *bytes.Buffer) *bufio.Reader {
	return bufio.NewReader(b) // указываем, от куда будет произведено чтение, из буфера (буфферизированный)
}

func main() {
	// Create a buffer for testing
	buffer := bytes.NewBufferString("Hello, World!")
	b := make([]byte, 13)
	r := getReader(buffer) // в данном случае лишний шаг т.к. buffer имеет метод Read()
	r.Read(b)
	fmt.Println(string(b)) // Hello, World!
}
