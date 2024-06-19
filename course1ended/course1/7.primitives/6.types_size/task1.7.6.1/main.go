package main

import (
	"fmt"
	"unsafe"
)

//Написать функцию для получения размера типа данных на языке Golang с помощью функции unsafe.Sizeof(). Необходимо написать функцию для каждого типа данных.

// sizeOfBool должна возвращать размер типа bool
func sizeOfBool(b bool) int {
	return int(unsafe.Sizeof(b))
}

// sizeOfInt должна возвращать размер типа int
func sizeOfInt(n int) int {
	return int(unsafe.Sizeof(n))
}

// sizeOfInt8 должна возвращать размер типа int8
func sizeOfInt8(n int8) int {
	return int(unsafe.Sizeof(n))
}

// — Функция  должна возвращать размер типа int16
func sizeOfInt16(n int16) int {
	return int(unsafe.Sizeof(n))
}

// — Функция  должна возвращать размер типа int32
func sizeOfInt32(n int32) int {
	return int(unsafe.Sizeof(n))
}

// — Функция  должна возвращать размер типа int64
func sizeOfInt64(n int64) int {
	return int(unsafe.Sizeof(n))
}

// — Функция олжна возвращать размер типа uint
func sizeOfUint(n uint) int {
	return int(unsafe.Sizeof(n))
}

// — Функция  должна возвращать размер типа uint8
func sizeOfUint8(n uint8) int {
	return int(unsafe.Sizeof(n))
}

func main() {

	var (
		bo        bool
		sizeInt   int
		sizeInt8  int8
		sizeInt16 int16
		sizeInt32 int32
		sizeInt64 int64
		sizeUint  uint
		sizeUint8 uint8
	)
	fmt.Printf("bool size: %d\n", sizeOfBool(bo))
	fmt.Printf("int size: %d\n", sizeOfInt(sizeInt))
	fmt.Printf("int8 size: %d\n", sizeOfInt8(sizeInt8))
	fmt.Printf("int16 size: %d\n", sizeOfInt16(sizeInt16))
	fmt.Printf("int32 size: %d\n", sizeOfInt32(sizeInt32))
	fmt.Printf("int64 size: %d\n", sizeOfInt64(sizeInt64))
	fmt.Printf("uint size: %d\n", sizeOfUint(sizeUint))
	fmt.Printf("uint8 size: %d\n", sizeOfUint8(sizeUint8))
}
