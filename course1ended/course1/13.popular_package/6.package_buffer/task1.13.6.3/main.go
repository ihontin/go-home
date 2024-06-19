package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func getDataString(b *bytes.Buffer) string {
	a, err := ioutil.ReadAll(b)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(b.String())
	return string(a)
}

//func main() {
//	// Create a new buffer
//	buffer := bytes.NewBufferString("Hello, World!")
//
//	// Call the getDataString function
//	result := getDataString(buffer)
//
//	// Check if the result matches the expected output
//	expected := "Hello, World!"
//	if result != expected {
//		panic(fmt.Sprintf("Expected %s, but got %s", expected, result))
//	}
//	fmt.Println(result)
//}
