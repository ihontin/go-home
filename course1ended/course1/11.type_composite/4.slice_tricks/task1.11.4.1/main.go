package main

// Cut (xs []int, start, end int) []int,
// которая будет вырезать элементы из среза xs начиная с индекса start и заканчивая индексом end.
// Функция должна возвращать новый срез, содержащий только вырезанные элементы.
func Cut(xs []int, start, end int) []int {
	lXs := len(xs) - 1
	var newList = make([]int, len(xs))
	if start < 0 || end < 0 || start > lXs || end+1 > lXs || len(xs) == 0 {
		return []int{}
	}
	_ = copy(newList, xs)
	return newList[start : end+1]
}

//func main() {
//	xs := []int{1, 2, 3, 4, 5}
//	result := Cut(xs, 1, 3)
//	fmt.Println(result) // Вывод: [2, 3, 4]
//}
