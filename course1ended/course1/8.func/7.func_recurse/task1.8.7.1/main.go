package main

import "fmt"

// Factorial (n int) int, которая вычисляет факториал числа n рекурсивно.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

//func rev(str string, str2 string) string {
//	if len(str2) == len(str) {
//		return str2
//	}
//	return rev(str, str2+(str[len(str)-(len(str2)+1):len(str)-len(str2)]))
//}

func main() {
	//a := "1234567"
	//fmt.Println(rev(a, ""))
	fmt.Println(Factorial(4))

}
