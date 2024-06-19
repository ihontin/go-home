package main

import (
	"fmt"
	"runtime"
	"time"
)

func factorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func factorialIterative(n int) int {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

// выдает true, если реализация быстрее и false, если медленнее
func compareWhichFactorialIsFaster() map[string]bool {
	var resMap = make(map[string]bool, 1)
	var facIterT, facRecurT time.Duration
	var res bool
	n := 100000

	start := time.Now()
	_ = factorialIterative(n)
	facIterT = time.Since(start)

	start = time.Now()
	_ = factorialRecursive(n)
	facRecurT = time.Since(start)

	if facIterT > facRecurT {
		res = true
	}
	//}
	resMap[fmt.Sprintf("%d", n)] = res
	return resMap
}

func main() {
	fmt.Println("Go version:", runtime.Version())                 //Версия go
	fmt.Println("Go OS/Arch:", runtime.GOOS, "/", runtime.GOARCH) // операционная система, архитектура

	fmt.Println("Which factorial is faster?")
	fmt.Println(compareWhichFactorialIsFaster())
}
