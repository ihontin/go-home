package main

import "fmt"

//Пример 1: односвязный список
//В этом примере мы создадим односвязный список и добавим несколько элементов в него.

//type Node struct {
//	value int
//	next  *Node
//}
//
//type LinkedList struct {
//	head *Node
//}
//
//func (list *LinkedList) Add(value int) {
//	newNode := &Node{value: value}
//
//	if list.head == nil {
//		list.head = newNode
//	} else {
//		current := list.head
//		for current.next != nil {
//			current = current.next
//		}
//		current.next = newNode
//	}
//}
//
//func main() {
//	list := LinkedList{}
//
//	list.Add(10)
//	list.Add(20)
//	list.Add(30)
//
//	current := list.head
//	for current != nil {
//		fmt.Println(current.value)
//		current = current.next
//	}
//}

//В этом примере мы создаем структуру Node, которая представляет узел списка с одним значением и ссылкой
//на следующий узел.Затем создаем структуру LinkedList, которая содержит указатель на голову списка.
//Метод Add добавляет новый узел в конец списка.В функции main мы создаем экземпляр связанного списка,
//добавляем несколько элементов и выводим значения элементов на экран.

//Пример 2: двусвязный список
//В этом примере мы реализуем двусвязный список и добавим несколько элементов в него.

//type Node struct {
//	value    int
//	previous *Node
//	next     *Node
//}
//
//type DoublyLinkedList struct {
//	head *Node
//}
//
//func (list *DoublyLinkedList) Add(value int) {
//	newNode := &Node{value: value}
//
//	if list.head == nil {
//		list.head = newNode
//	} else {
//		current := list.head
//		for current.next != nil {
//			current = current.next
//		}
//		current.next = newNode
//		newNode.previous = current
//	}
//}
//
//func main() {
//	list := DoublyLinkedList{}
//
//	list.Add(10)
//	list.Add(20)
//	list.Add(30)
//
//	current := list.head
//	for current != nil {
//		fmt.Println(current.value)
//		current = current.next
//	}
//}

//В этом примере мы создаем структуру Node, которая представляет узел списка с одним значением,
//ссылкой на предыдущий узел и ссылкой на следующий узел.Затем мы создаем структуру DoublyLinkedList,
//которая содержит указатель на голову списка.Метод Add добавляет новый узел в конец списка и устанавливает ссылку
//на предыдущий узел.В функции main мы создаем экземпляр двусвязного списка,
//добавляем несколько элементов и выводим значения элементов на экран.

//Пример 3: кольцевой список
//В этом примере мы реализуем кольцевой список и добавим несколько элементов в него.

type Node struct {
	value int
	next  *Node
}

type CircularLinkedList struct {
	head *Node
}

func (list *CircularLinkedList) Add(value int) {
	//newNode := &Node{value: value}
	//
	//if list.head == nil {
	//	list.head = newNode
	//	newNode.next = newNode
	//} else {
	//	current := list.head
	//	for current.next != list.head {
	//		current = current.next
	//	}
	//	current.next = newNode
	//	newNode.next = list.head
	//}
	newN := &Node{value: value}

	if list.head == nil {
		list.head = newN
		newN.next = newN
	} else {
		current := list.head
		for current != list.head {
			current = current.next
		}
		current.next = newN
		newN.next = list.head
	}
}

func main() {
	list := CircularLinkedList{}

	list.Add(10)
	list.Add(20)
	list.Add(30)

	current := list.head
	for {
		fmt.Println(current.value)
		current = current.next
		if current == list.head {
			break
		}
	}
}

//В этом примере мы создаем структуру Node, которая представляет узел списка с одним значением и ссылкой
//на следующий узел.Затем мы создаем структуру CircularLinkedList, которая содержит указатель на голову списка.
//Метод Add добавляет новый узел в конец списка и устанавливает ссылку на голову списка для последнего узла.
//В функции main мы создаем экземпляр кольцевого списка, добавляем несколько элементов и выводим значения элементов на экран.
