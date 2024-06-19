package main

//RemoveExtraMemory,которая удаляет лишний объем памяти из среза xs типа []int.

func RemoveExtraMemory(xs []int) []int {
	var unicMap = make(map[int]bool)
	for _, val := range xs {
		if !unicMap[val] {
			unicMap[val] = true
		}
	}
	var newList = make([]int, 0, len(unicMap))
	for _, val := range xs {
		if unicMap[val] {
			delete(unicMap, val)
			newList = append(newList, val)
		}
	}
	return newList
}

//func main() {
//	var xs []int
//	xs = RemoveExtraMemory(xs)
//	fmt.Println(cap(xs), xs)
//	xs = append(xs, 2, 3, 4, 2, 3, 4, 5, 6, 7)
//	xs = RemoveExtraMemory(xs)
//	fmt.Println(cap(xs), xs, "finish")
//}
