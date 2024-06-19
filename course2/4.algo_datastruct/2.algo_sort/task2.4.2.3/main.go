package main

// Структура пользователя
type User struct {
	ID   int
	Name string
	Age  int
}

// Функция слияния двух отсортированных массивов пользователей
func Merge(arr1 []User, arr2 []User) []User {
	if len(arr1) < 1 && len(arr2) < 1 {
		return []User{}
	} else if len(arr1) < 1 {
		return arr2
	} else if len(arr2) < 1 {
		return arr1
	}

	merged := make([]User, 0, len(arr1)+len(arr2))

	index1, index2 := 0, 0
	for index1 < len(arr1) && index2 < len(arr2) {
		if arr1[index1].ID < arr2[index2].ID {
			merged = append(merged, arr1[index1])
			index1++
		} else {
			merged = append(merged, arr2[index2])
			index2++
		}
	}
	for ; index1 < len(arr1); index1++ {
		merged = append(merged, arr1[index1])
	}
	for ; index2 < len(arr2); index2++ {
		merged = append(merged, arr2[index2])
	}

	return merged
}
