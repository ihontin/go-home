package main

import (
	"fmt"
)

type User struct {
	ID int
}

// ----------------------------------------- пузырьком сортировка
//func buble(u []User) {
//	for i := 0; i < len(u)-1; i++ {
//		for j := i + 1; j < len(u); j++ {
//			if u[i].ID > u[j].ID {
//				u[i].ID, u[j].ID = u[j].ID, u[i].ID
//			}
//		}
//	}
//}
//func main() {
//	user := []User{User{5}, User{1}, User{22}, User{14}, User{9}, User{5}}
//	buble(user)
//	fmt.Println(user)
//}

// ----------------------------------------- сортировка выбором

//func selection(u []User) {
//	for i := 0; i < len(u); i++ {
//		minIndex := i
//		for j := i + 1; j < len(u); j++ {
//			if u[j].ID < u[minIndex].ID {
//				minIndex = j
//			}
//		}
//		u[i], u[minIndex] = u[minIndex], u[i]
//	}
//}

//func main() {
//	user := []User{User{5}, User{1}, User{22}, User{14}, User{9}, User{5}}
//	selection(user)
//	fmt.Println(user)
//}

// ----------------------------------------- сортировка вставками

//func insertionSort(u []User) {
//	for i := 1; i < len(u); i++ {
//		key := u[i]
//		j := i - 1
//		for j >= 0 && u[j].ID > key.ID {
//			u[j+1] = u[j]
//			j--
//		}
//		u[j+1] = key
//	}
//}
//func main() {
//	user := []User{User{5}, User{1}, User{22}, User{14}, User{9}, User{5}}
//	insertionSort(user)
//	fmt.Println(user)
//}

// -----------------------------------------быстрая сортировка рекурсией

//func pivot(u []User, low, high int) int {
//	hiU := u[high]
//	l := low
//	for i := low; i < high; i++ {
//		if u[i].ID < hiU.ID {
//			u[l], u[i] = u[i], u[l]
//			l++
//		}
//	}
//	u[l], u[high] = u[high], u[l]
//	return l
//}
//
//func quick(u []User, low, high int) {
//	if low < high {
//		pi := pivot(u, low, high)
//		quick(u, low, pi-1)
//		quick(u, pi+1, high)
//	}
//}
//
//func q(u []User) {
//	quick(u, 0, len(u)-1)
//}
//
//func main() {
//	user := []User{User{5}, User{1}, User{22}, User{14}, User{9}, User{5}}
//	q(user)
//	fmt.Println(user)
//}

// -----------------------------------------быстрая сортировка

func quickSort(u []User) []User {
	if len(u) < 2 {
		return u
	}
	piv := u[0]
	var less, greater []User
	for _, num := range u[1:] {
		if num.ID <= piv.ID {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	res := append(append(quickSort(less), piv), quickSort(greater)...)
	return res
}

func quick(u []User) {
	for i, val := range quickSort(u) {
		u[i] = val
	}
}

func main() {
	user := []User{User{5}, User{1}, User{22}, User{14}, User{9}, User{5}}
	quick(user)
	fmt.Println(user)
}

// ------------------------------------------ сортировка слиянием

//func mergeSort(u []User) []User {
//	if len(u) < 2 {
//		return u
//	}
//	mid := len(u) / 2
//	left := mergeSort(u[:mid])
//	right := mergeSort(u[mid:])
//	return merge(left, right)
//}
//
//func merge(l, r []User) []User {
//	var merged []User
//	for len(l) > 0 && len(r) > 0 {
//		if l[0].ID <= r[0].ID {
//			merged = append(merged, l[0])
//			l = l[1:]
//		} else {
//			merged = append(merged, r[0])
//			r = r[1:]
//		}
//	}
//	merged = append(append(merged, l...), r...)
//	return merged
//}
//
//func merging(u []User) {
//	for i, val := range mergeSort(u) {
//		u[i] = val
//	}
//}
//func main() {
//	user := []User{User{5}, User{1}, User{22}, User{14}, User{9}, User{5}}
//	merging(user)
//	fmt.Println(user)
//}
