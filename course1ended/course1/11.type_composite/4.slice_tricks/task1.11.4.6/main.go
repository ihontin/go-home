package main

// InsertToStart , которая принимает срез xs типа []int и произвольное количество аргументов типа int.
// Функция должна вставить переданные аргументы в начало среза xs и вернуть получившийся срез.
func InsertToStart(xs []int, x ...int) []int {
	lenX := len(x)
	var newList = make([]int, len(xs), lenX+len(xs))
	_ = copy(newList, xs)
	newList = append(x, xs...)
	return newList
}

//func main() {
//	xs := []int{1, 2, 3}
//	result := InsertToStart(xs)
//	xs[0] = 22
//	fmt.Println(result, xs)
//	result должен быть равен []int{4, 5, 6, 1, 2, 3}
//}
