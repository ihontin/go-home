package main

import "fmt"

func main() {
	var num = 1234

	sum := num%100 + num/100
	fmt.Println(sum, num%100, num/100)
}
