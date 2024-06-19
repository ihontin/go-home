package main

import "fmt"

//Operate, которая принимает произвольное количество аргументов типа interface{}
//и функцию func(xs ...interface{}) interface{}.
//Функция Operate должна применять переданную функцию к аргументам и возвращать результат.

var Operate func(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} = funcOperate // реализуй меня
func funcOperate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

var Concat func(xs ...interface{}) interface{} = funcConcat // реализуй меня для string
func funcConcat(xs ...interface{}) interface{} {
	var returnStr string
	for _, x := range xs {
		switch x.(type) {
		case string:
			returnStr += x.(string)
		default:
			return nil
		}
	}
	return returnStr
}

var Sum func(xs ...interface{}) interface{} = funcSum // реализуй меня для int и float64

func funcSum(xs ...interface{}) interface{} {
	var returnInt int
	var returnFloat float64
	for i, x := range xs {
		switch x.(type) {
		case int:
			returnInt += x.(int)
			if i == len(xs)-1 {
				return returnInt
			}
		case float64:
			returnFloat += x.(float64)
			if i == len(xs)-1 {
				return returnFloat
			}
		default:
			return nil
		}
	}
	return nil
}

func main() {
	fmt.Println(Operate(Concat, "Hello, ", "World!"))  // Вывод: "Hello, World!"
	fmt.Println(Operate(Sum, 1, 2, 3, 4, 5))           // Вывод: 15
	fmt.Println(Operate(Sum, 1.1, 2.2, 3.3, 4.4, 5.5)) // Вывод: 16.5
}
