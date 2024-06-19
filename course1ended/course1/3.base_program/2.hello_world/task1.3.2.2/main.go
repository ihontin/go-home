package main

import "fmt"

// —  Программа должна вывести первой надписью Hello world!
func HelloWorld() string {
	return "Hello world!"
}

// — Второй надписью должна быть This is second line!
func SecondString() string {
	return "This is second line!"
}

// —  Третьей надписью должна быть This is third line!
func ThirdString() string {
	return "This is third line!"
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(SecondString())
	fmt.Println(ThirdString())
}
