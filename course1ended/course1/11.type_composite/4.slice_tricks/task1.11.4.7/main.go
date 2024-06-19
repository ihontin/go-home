package main

import "fmt"

// Pop , которая принимает на вход срез xs типа []int и возвращает первый элемент этого среза,
// а также новый срез, в котором удален первый элемент.
func Pop(xs []int) (int, []int) {
	lenXs := len(xs)
	if lenXs == 0 {
		return 0, []int{}
	}
	var newList = make([]int, 0, lenXs-1)
	var firstEl = xs[0]
	newList = xs[1:]
	return firstEl, newList
}

func main() {
	var val int
	var list, newL = []int{3, 6}, make([]int, 0, 5)
	val, newL = Pop(list)
	fmt.Printf("Значение: %d, Новый срез: %v", val, newL)
}
