package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

//— Функция должна корректно конвертировать двоичную строку в число с плавающей точкой типа float32

//— Функция должна вернуть число 0.15625 для строки “00111110001000000000000000000000”

func binaryStringToFloat(binary string) float32 {
	var number uint32
	// Преобразование строки в двоичной системе в целочисленное представление
	number1, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}
	number = *(*uint32)(unsafe.Pointer(&number1))
	// Преобразование целочисленного представления в число с плавающей точкой
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}

func main() {
	fmt.Println(binaryStringToFloat("00111110001000000000000000000000"))
}
