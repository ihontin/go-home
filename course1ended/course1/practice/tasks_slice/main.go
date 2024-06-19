package main

import "fmt"

func main() {
	//var a = []int{1, 3, 8, 4, 9, 9, 8, 8}
	//var b = []int{1, 2, 3, 2, 4}
	//fmt.Println("start", a, b)
	//c := delDubl(a, b)
	//fmt.Println("start", a, b)
	//fmt.Println("finish", c)

	var a = []int{1, 3, 8, 4, 9, 9, 8, 8}
	fmt.Println(a)
	b := div2(a)
	fmt.Println(b)
}

func div2(l []int) []int {
	var outSlice = make([]int, 0, len(l))
	for _, val := range l {
		if val%2 == 0 {
			outSlice = append(outSlice, val)
		}
	}
	return outSlice
}

//func delDubl(l1, l2 []int) []int {
//	l1 = append(l1, l2...)
//	newLen := len(l1) + len(l2)
//	var uniqMap = make(map[int]bool, newLen)
//	var outSlice = make([]int, 0, newLen)
//	for _, val := range l1 {
//		if _, ok := uniqMap[val]; !ok {
//			uniqMap[val] = true
//			outSlice = append(outSlice, val)
//		}
//	}
//	return outSlice
//}
