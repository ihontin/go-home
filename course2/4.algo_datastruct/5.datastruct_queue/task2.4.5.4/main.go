package main

// Пример кода на языке программирования

// Определение стека
var stack []int

// Функция для добавления элемента в стек
func push(value int) {
	stack = append(stack, value)
}

// Функция для удаления и возврата последнего элемента из стека
func pop() int {
	if len(stack) == 0 {
		panic("Стек пуст")
	}
	lastIndex := len(stack) - 1
	value := stack[lastIndex]
	stack = stack[:lastIndex]
	return value
}

func main() {
	// Пример использования стека для операций
	push(5)
	push(3)
	result := pop() + pop()
	push(result)
}
