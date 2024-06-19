package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	myPanic(ch)
	fmt.Println(<-ch)
}

func myPanic(ch chan string) {
	defer func() {
		if r := recover(); r != nil {
			ch <- fmt.Sprintf("%v", r)
		}
	}()
	panic("my panic message")
}
