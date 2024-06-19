package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Задача 1
// Получение максимального значения для типа
// Если значение не нулевое, вернуть максимальное значение для типа
func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {
	if in8 != 0 {
		in8 = 1<<(getBits(in8)-1) - 1
	}
	if in16 != 0 {
		in16 = 1<<(getBits(in16)-1) - 1
	}
	if in32 != 0 {
		in32 = 1<<(getBits(in32)-1) - 1
	}
	if in64 != 0 {
		in64 = 1<<(getBits(in64)-1) - 1
	}
	return in8, in16, in32, in64
}

// Если значение не нулевое, вернуть максимальное значение для типа
func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	if uin8 != 0 {
		uin8 = 1<<getBits(uin8) - 1

	}
	if uin16 != 0 {
		uin16 = 1<<getBits(uin16) - 1

	}
	if uin32 != 0 {
		uin32 = 1<<getBits(uin32) - 1

	}
	if uin64 != 0 {
		uin64 = 1<<getBits(uin64) - 1

	}
	return uin8, uin16, uin32, uin64
}

func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)

	return bits
}

// Задача 2
// Функция при передаче одинаковых переменных возвращает true используя bitwise операции
func isEquals(a, b int) bool {
	return a^b == 0
}

// Задача 3
// Написать функцию которая принимает указатель на переменную и увеличивает ее значение на 1, с помощью инкремента
func incrementPointer(ptr *int) {
	*ptr++
}
func main() {
	//1
	fmt.Println(getIntMaxValue(1, 1, 1, 1))
	fmt.Println(getUintMaxValue(1, 1, 1, 1))
	//2
	var (
		x = 20
		y = 20
	)
	fmt.Println(isEquals(x, y))
	// 3
	var a = 86
	fmt.Println(a)
	incrementPointer(&a)
	fmt.Println(a)
}
