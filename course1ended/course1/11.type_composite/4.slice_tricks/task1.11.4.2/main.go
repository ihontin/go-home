package main

// InsertAfterIDX (xs []int, idx int, x... int) []int,
// которая вставляет элементы x в массив xs после указанного индекса idx.
func InsertAfterIDX(xs []int, idx int, x ...int) []int {
	lXs := len(xs) - 1
	if idx < 0 || idx > lXs || len(xs) == 0 {
		return []int{}
	}
	var newList = make([]int, lXs+len(x))
	newList = append(xs[:idx+1], append(x, xs[idx+1:]...)...)
	return newList
}

//func main() {
//	xs := []int{1, 2, 3, 4, 5}
//	result := InsertAfterIDX(xs, 4, 6, 7, 8)
//	fmt.Println(result) // Вывод: [1 2 3 6 7 8 4 5]
//}
