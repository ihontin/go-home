package main

// Shift принимает на вход срез целых чисел и выполняет операцию сдвига элементов вправо на одну позицию.
// Функция должна возвращать первый элемент сдвинутого среза и сам сдвинутый срез.
func Shift(xs []int) (int, []int) {
	lenXs := len(xs)
	// проверка на пустой список
	if lenXs == 0 {
		return 0, []int{}
	}
	var newList = make([]int, 0, lenXs)        // создание нового среза с емкостью принятого среза
	var firstEl = xs[0]                        // инициализация переменной, что будет хранить значение первого элемента
	newList = append(newList, xs[lenXs-1])     // добавление последнего элемента списка xs в начало нового среза
	newList = append(newList, xs[:lenXs-1]...) // добавление оставшихся значений от 0 до предпоследнего в новый список
	return firstEl, newList
}

//func main() {
//	xs := []int{1, 2, 3, 4, 5}
//	firstElement, shiftedSlice := Shift(xs)
//	fmt.Println(firstElement) // Вывод: 1
//	fmt.Println(shiftedSlice) // Вывод: [5 1 2 3 4]
//}
