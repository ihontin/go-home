package main

import (
	"fmt"
	"os"
)

func main() {
	a, b := 8, 13
	fmt.Println(*testDefer(&a, &b))
	file, err := os.OpenFile("new.txt", os.O_RDWR|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	err = writeToFile(file, "Hello World!")
	if err != nil {
		fmt.Println(err)
	}
}
func writeToFile(file *os.File, data string) error {
	_, err := file.Write([]byte(data))
	return err
}

func testDefer(a, b *int) *int {
	var c int
	defer func() {

	}()
	c = sum(*a, *b)
	return &c
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}
