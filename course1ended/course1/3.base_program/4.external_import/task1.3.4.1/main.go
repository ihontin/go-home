package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

// DecimalSum функция, которая принимает две строки, содержащие числа с плавающей точкой,
// складывает их и возвращает результат в виде строки.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalSum(a, b string) (string, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	num2, err := decimal.NewFromString(b)
	checkErr(err)
	res := num1.Add(num2)
	return res.String(), err
}

// DecimalSubtract функция, которая принимает две строки, содержащие числа с плавающей точкой,
// вычитает из первого числа второе и возвращает результат в виде строки.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalSubtract(a, b string) (string, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	num2, err := decimal.NewFromString(b)
	checkErr(err)
	res := num1.Sub(num2)
	return res.String(), err
}

// DecimalMultiply функция, которая принимает две строки, содержащие числа с плавающей точкой,
// перемножает их и возвращает результат в виде строки.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalMultiply(a, b string) (string, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	num2, err := decimal.NewFromString(b)
	checkErr(err)
	res := num1.Mul(num2)
	return res.String(), err
}

// DecimalDivide функция, которая принимает две строки, содержащие числа с плавающей точкой,
// делит первое число на второе и возвращает результат в виде строки.
// Если входные данные некорректны или происходит на ноль, функция должна возвращать ошибку.
func DecimalDivide(a, b string) (string, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	num2, err := decimal.NewFromString(b)
	checkErr(err)
	res := num1.Div(num2)
	return res.String(), err
}

// DecimalRound функция, которая принимает строку, содержащую число с плавающей точкой,
// и точность округления в виде int32. Функция должна округлить число до указанной точности и вернуть результат в виде строки.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalRound(a string, precision int32) (string, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	res := num1.Round(precision)
	return res.String(), err
}

// DecimalGreaterThan — функция, которая принимает две строки, содержащие числа с плавающей точкой,
// и возвращает true, если первое число больше второго, и false в противном случае.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalGreaterThan(a, b string) (bool, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	num2, err := decimal.NewFromString(b)
	checkErr(err)
	res := num1.GreaterThan(num2)
	return res, err
}

// DecimalLessThan функция, которая принимает две строки, содержащие числа с плавающей точкой,
// и возвращает true, если первое число меньше второго, и false в противном случае.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalLessThan(a, b string) (bool, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	num2, err := decimal.NewFromString(b)
	checkErr(err)
	res := num1.LessThan(num2)
	return res, err
}

// DecimalEqual функция, которая принимает две строки, содержащие числа с плавающей точкой,
// и возвращает true, если числа равны, и false в противном случае.
// Если входные данные некорректны, функция должна возвращать ошибку.
func DecimalEqual(a, b string) (bool, error) {
	num1, err := decimal.NewFromString(a)
	checkErr(err)
	num2, err := decimal.NewFromString(b)
	checkErr(err)
	res := num1.Equal(num2)
	return res, err
}

func main() {
	var a string
	var b string
	var p int32

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&p)

	resultSum, err := DecimalSum(a, b)
	checkErr(err)
	fmt.Println(resultSum)

	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	resultSub, err := DecimalSubtract(a, b)
	checkErr(err)
	fmt.Println(resultSub)

	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	resultMul, err := DecimalMultiply(a, b)
	checkErr(err)
	fmt.Println(resultMul)

	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	resultDiv, err := DecimalDivide(a, b)
	checkErr(err)
	fmt.Println(resultDiv)

	//fmt.Scanln(&a)
	//fmt.Scanln(&p)
	resultRou, err := DecimalRound(a, p)
	checkErr(err)
	fmt.Println(resultRou)

	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	resultGre, err := DecimalGreaterThan(a, b)
	checkErr(err)
	fmt.Println(resultGre)

	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	resultLes, err := DecimalLessThan(a, b)
	checkErr(err)
	fmt.Println(resultLes)

	//fmt.Scanln(&a)
	//fmt.Scanln(&b)
	resultEqu, err := DecimalEqual(a, b)
	checkErr(err)
	fmt.Println(resultEqu)
}
