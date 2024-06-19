package main

// RemoveIDX (xs []int, idx int) []int, которая удаляет элемент из среза xs
// по указанному индексу idx и возвращает новый срез без этого элемента.
func RemoveIDX(xs []int, idx int) []int {
	lXs := len(xs)
	if lXs == 0 {
		return []int{}
	}
	if idx < 0 || idx > lXs-1 {
		return xs
	}
	xs = append(xs[:idx], xs[idx+1:]...)
	return xs
}
