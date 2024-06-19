package main

//Необходимо реализовать функцию appendInt, которая принимает срез xs типа []int и
//произвольное количество аргументов типа int, и возвращает новый срез []int,
//в котором все аргументы добавлены в конец исходного среза.

func appendInt(xs []int, x ...int) []int {
	return append(xs, x...)
}

//func main() {
//	var a = []int{1, 2, 3, 4}
//	a = appendInt(a, 0, 9, 8, 7)
//	fmt.Println(a)
//}
