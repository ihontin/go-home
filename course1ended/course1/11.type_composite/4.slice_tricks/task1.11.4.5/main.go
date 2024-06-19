package main

//FilterDividers (xs []int, divider int) []int, которая фильтрует элементы входного среза xs
//и возвращает новый срез, содержащий только те числа, которые делятся на divider без остатка.

func FilterDividers(xs []int, divider int) []int {
	lXs := len(xs)
	if lXs == 0 {
		return []int{}
	}
	if divider == 0 {
		return xs
	}
	var newList = make([]int, 0, lXs)
	var dVal int
	for _, num := range xs {
		if num%divider == 0 {
			dVal = num
			newList = append(newList, dVal)
		}
	}
	return newList
}

//func main() {
//	var xs = []int{2, 3, 4, 2, 3, 4, 5, 6, 7}
//	ys := FilterDividers(xs, 7)
//	fmt.Println(xs, ys)
//}
