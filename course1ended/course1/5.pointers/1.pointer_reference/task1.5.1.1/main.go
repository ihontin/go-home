package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Add должна принимать два целых числа и возвращать указатель на результат их сложения.
func Add(a, b int) *int {
	res := a + b
	return &res
}

// Max должна принимать срез целых чисел и возвращать указатель на максимальное число в этом срезе.
func Max(numbers []int) *int {
	var maxNum int
	for i, _ := range numbers {
		if numbers[i] > numbers[maxNum] {
			maxNum = i
		}
	}
	return &numbers[maxNum]
}

// IsPrime должна принимать целое число и возвращать указатель на логическое значение, которое указывает, является ли число простым.
func IsPrime(number int) *bool {
	simple := true
	notSimple := false
	if number < 2 {
		return &notSimple
	} else if number == 2 {
		return &simple
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return &notSimple
		}
	}
	return &simple
}

// ConcatenateStrings должна принимать срез строк и возвращать указатель на строку, которая является результатом их конкатенации.
func ConcatenateStrings(strs []string) *string {
	var concString string
	for _, s := range strs {
		concString += s
	}
	return &concString
}

func main() {
	//1 task
	var a, b int
	fmt.Scanln(&a, &b)
	Add(a, b)

	//2 task
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	sliceInput := strings.Split(input, " ")

	var sliceInt = make([]int, 0, len(sliceInput))
	for _, part := range sliceInput {
		s, err2 := strconv.Atoi(part)
		if err2 != nil {
			panic(err2)
		}
		sliceInt = append(sliceInt, s)
	}
	Max(sliceInt)

	//3 task
	var c int
	fmt.Scan(&c)
	IsPrime(c)

	//4 task
	inputNewStr, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}
	inputNewStr = strings.TrimSuffix(inputNewStr, "\n")
	inputNewStr = strings.TrimSpace(inputNewStr)
	sliceString := strings.Split(inputNewStr, " ")
	ConcatenateStrings(sliceString)
}
