package main

//Пример 1: сортировка пузырьком
//Рассмотрим алгоритм сортировки пузырьком, который имеет временную сложность O(n^2).

//func bubbleSort(arr []int) {
//	n := len(arr)
//	for i := 0; i < n-1; i++ {
//		for j := 0; j < n-i-1; j++ {
//			if arr[j] > arr[j+1] {
//				arr[j], arr[j+1] = arr[j+1], arr[j]
//			}
//		}
//	}
//}
//
//func main() {
//	arr := []int{64, 34, 25, 12, 22, 11, 90}
//	bubbleSort(arr)
//	fmt.Println("Sorted array:", arr)
//}

//Пример 2: сортировка выбором
//Рассмотрим алгоритм сортировки выбором, который имеет временную сложность O(n^2).

//func selectionSort(arr []int) {
//	n := len(arr)
//	for i := 0; i < n-1; i++ {
//		minIndex := i
//
//		for j := i + 1; j < n; j++ {
//			if arr[j] < arr[minIndex] {
//				minIndex = j
//			}
//		}
//		arr[i], arr[minIndex] = arr[minIndex], arr[i]
//	}
//}
//
//func main() {
//	arr := []int{64, 34, 25, 12, 22, 14, 90}
//	selectionSort(arr)
//	fmt.Println("Sorted array:", arr)
//}

//Пример 3: сортировка вставками
//Рассмотрим алгоритм сортировки вставками, который имеет временную сложность O(n^2).

//func insertionSort(arr []int) {
//	n := len(arr)
//	for i := 1; i < n; i++ {
//		key := arr[i]
//		j := i - 1
//		for j >= 0 && arr[j] > key {
//			arr[j+1] = arr[j]
//			j--
//		}
//		arr[j+1] = key
//	}
//}
//
//func main() {
//	arr := []int{64, 34, 25, 12, 22, 11, 90}
//	insertionSort(arr)
//	fmt.Println("Sorted array:", arr)
//}

//Пример 4: быстрая сортировка
//Рассмотрим алгоритм быстрой сортировки, который имеет временную сложность O(n log n).

//func quickSort(arr []int, low, high int) {
//	if low < high {
//		pi := partition(arr, low, high)
//		quickSort(arr, low, pi-1)
//		quickSort(arr, pi+1, high)
//	}
//}
//
//func partition(arr []int, low, high int) int {
//	pivot := arr[high]
//	i := low - 1
//	for j := low; j < high; j++ {
//		if arr[j] < pivot {
//			i++
//			arr[i], arr[j] = arr[j], arr[i]
//		}
//	}
//	arr[i+1], arr[high] = arr[high], arr[i+1]
//	return i + 1
//}
//
//func main() {
//	arr := []int{64, 34, 25, 12, 22, 11, 90}
//	n := len(arr)
//	quickSort(arr, 0, n-1)
//	fmt.Println("Sorted array:", arr)
//}

//Пример 5: сортировка слиянием
//Рассмотрим алгоритм сортировки слиянием, который имеет пространственную сложность O(n).

//func mergeSort(arr []int) []int {
//	if len(arr) <= 1 {
//		return arr
//	}
//	mid := len(arr) / 2
//	left := mergeSort(arr[:mid])
//	right := mergeSort(arr[mid:])
//	return merge(left, right)
//}
//
//func merge(left, right []int) []int {
//	result := make([]int, 0, len(left)+len(right))
//	for len(left) > 0 && len(right) > 0 {
//		if left[0] <= right[0] {
//			result = append(result, left[0])
//			left = left[1:]
//		} else {
//			result = append(result, right[0])
//			right = right[1:]
//		}
//	}
//	result = append(result, left...)
//	result = append(result, right...)
//	return result
//}
//
//func main() {
//	arr := []int{64, 34, 25, 12, 22, 11, 90}
//	sortedArr := mergeSort(arr)
//	fmt.Println("Sorted array:", sortedArr)
//}

//Пример 6: сортировка подсчетом
//Рассмотрим алгоритм сортировки подсчетом, который имеет пространственную сложность O(n).

//func countingSort(arr []int) {
//	max := findMax(arr)
//	count := make([]int, max+1)
//	sortedArr := make([]int, len(arr))
//
//	for _, num := range arr {
//		count[num]++
//	}
//
//	for i := 1; i <= max; i++ {
//		count[i] += count[i-1]
//	}
//
//	for i := len(arr) - 1; i >= 0; i-- {
//		num := arr[i]
//		sortedArr[count[num]-1] = num
//		count[num]--
//	}
//
//	for i := 0; i < len(arr); i++ {
//		arr[i] = sortedArr[i]
//	}
//}
//
//func findMax(arr []int) int {
//	max := arr[0]
//	for _, num := range arr {
//		if num > max {
//			max = num
//		}
//	}
//	return max
//}
//
//func main() {
//	arr := []int{64, 34, 25, 12, 22, 11, 90}
//	countingSort(arr)
//	fmt.Println("Sorted array:", arr)
//}

//Пример 7: сортировка кучей
//Рассмотрим алгоритм сортировки кучей, который имеет пространственную сложность O(n).

//func heapSort(arr []int) {
//	n := len(arr)
//
//	for i := n/2 - 1; i >= 0; i-- {
//		heapify(arr, n, i)
//	}
//
//	for i := n - 1; i > 0; i-- {
//		arr[0], arr[i] = arr[i], arr[0]
//		heapify(arr, i, 0)
//	}
//}
//
//func heapify(arr []int, n, i int) {
//	largest := i
//	left := 2*i + 1
//	right := 2*i + 2
//
//	if left < n && arr[left] > arr[largest] {
//		largest = left
//	}
//
//	if right < n && arr[right] > arr[largest] {
//		largest = right
//	}
//
//	if largest != i {
//		arr[i], arr[largest] = arr[largest], arr[i]
//		heapify(arr, n, largest)
//	}
//}
//
//func main() {
//	arr := []int{64, 34, 25, 12, 22, 11, 90}
//	heapSort(arr)
//	fmt.Println("Sorted array:", arr)
//}
