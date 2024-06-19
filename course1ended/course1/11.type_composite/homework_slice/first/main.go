package main

import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := NewSlice(3, 2)
	fmt.Println(s)
	s = Append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s)
	ss := make([]byte, 100)
	for i := range ss {
		ss[i] = 1
	}

	//ptr := C.malloc(100)
	//defer C.free(ptr)
	//fmt.Println(ptr)
}

// NewSlice - create a new slice
func NewSlice(opts ...int) []byte {
	if opts == nil {
		return nil
	}

	capacity := 0
	length := 0
	if len(opts) == 1 {
		capacity = opts[0]
		length = opts[0]
	}
	if len(opts) == 2 {
		length = opts[0]
		capacity = opts[1]
	}

	if length > capacity {
		capacity = length
	}

	// здесь мы создаем slice header
	sh := &reflect.SliceHeader{
		Len: length,
		Cap: capacity,
	}

	// здесь мы выделяем память для хранения элементов слайса
	for i := 0; i < capacity; i++ {
		b := new(byte)
		if i == 0 {
			sh.Data = uintptr(unsafe.Pointer(b))
		}
	}

	return *(*[]byte)(unsafe.Pointer(sh))
}

func Append(slice []byte, data ...byte) []byte {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	if len(data) == 0 {
		return slice
	}

	idx := sh.Len

	if sh.Cap < sh.Len+len(data) {
		oldSlice := slice
		slice = NewSlice(sh.Len+len(data), (sh.Cap+len(data))*2)
		for i := range oldSlice {
			slice[i] = oldSlice[i]
		}

		sh = (*reflect.SliceHeader)(unsafe.Pointer(&oldSlice))
	}

	for i := range data {
		slice[idx] = data[i]
		idx++
	}

	return slice
}
