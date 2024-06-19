package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type LinkedLister interface {
	LoadData(path string) error      //
	Init(c []Commit)                 //
	Len() int                        //
	SetCurrent(n int) error          //
	Current() *Node                  //
	Next() *Node                     //
	Prev() *Node                     //
	Insert(n int, c Commit) error    //
	Push(c Commit) error             //
	Delete(n int) error              //
	DeleteCurrent() error            //
	Index() (int, error)             //
	GetByIndex(n int) (*Node, error) //
	Pop() *Node                      //
	Shift() *Node                    //
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

func (d *DoubleLinkedList) SetCurrent(n int) error {
	if n < 0 || n >= d.len {
		return fmt.Errorf("index out of range")
	}

	node := d.head
	for i := 0; i < n; i++ {
		node = node.next
	}

	d.curr = node
	return nil
}

func (d *DoubleLinkedList) Init(c []Commit) {
	for _, commit := range c {
		err := d.Push(commit)
		if err != nil {
			log.Fatal("error in Push method:", err)
		}
	}
}

func quicksort(arr []Commit, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func partition(arr []Commit, low, high int) int {
	if len(arr) < 2 {
		return 0
	}
	pivot := int(arr[high].Date.Unix())
	i := low - 1
	for j := low; j < high; j++ {
		if int(arr[j].Date.Unix()) < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// LoadData загрузка данных из подготовленного json файла
func (d *DoubleLinkedList) LoadData(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var commitList []Commit
	err = json.NewDecoder(file).Decode(&commitList)

	quicksort(commitList, 0, len(commitList)-1) // отсортировать список используя самописный QuickSort
	d.Init(commitList)
	return nil
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	return d.curr
}

// Prev получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {
	if d.curr != nil && d.curr.next != nil {
		d.curr = d.curr.next
	}
	return d.curr
}

// Insert вставка элемента после n элемента
func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return fmt.Errorf("index out of range")
	}
	newNode := &Node{data: &c} // создаем новый узел
	if n == 0 {                // если нужно добавить в начало
		newNode.next = d.head // связываем новый узел со старым списком
		if d.head != nil {    // если в старом списке нет элементов
			d.head.prev = newNode // предыдущему узлу даем ссылку на текущий
		}
		d.head = newNode // новый узел делаем первым
	} else { // если нужно добавить НЕ в начало
		prev := d.head // ссылку на голову копируем в новую переменную

		for i := 1; i < n; i++ {
			if prev == nil { // проверка в цикле на выход за границы скписка
				return fmt.Errorf("index out of range")
			}
			prev = prev.next // проходим по списку до нужного индекса вставки n
		}
		// вставляем новое значение в список
		newNode.next = prev.next // новому узлу добавляем ссылку на следующий элемент списка начиная с n индекса
		newNode.prev = prev      // новому узлу добавляем ссылку на предыдущий элемент списка по n индексу
		if prev.next != nil {    // если текущий элемент списка НЕ последний
			prev.next.prev = newNode // следующий узел связываем с новым узлом, добавляя ссылку на новый узел
		} else {
			d.tail = newNode // если текущий узел списка последний, он становится хвостом
		}
		prev.next = newNode // новый узел "newNode" по индексу n теперь является prev.next
	}
	d.len++ // инкремент длинны списка
	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	if n < 0 || n >= d.len {
		return fmt.Errorf("index out of range")
	}
	var node *Node
	if n == 0 {
		node = d.head
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		} else {
			d.tail = nil
		}
	} else {
		node = d.head
		for i := 0; i < n; i++ {
			if node == nil {
				return fmt.Errorf("index out of range")
			}
			node = node.next
		}
		if node.prev != nil {
			node.prev.next = node.next
		} else {
			d.head = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		} else {
			d.tail = node.prev
		}
	}
	d.len--
	return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.curr == nil {
		return fmt.Errorf("no current node")
	}

	if d.curr.prev != nil {
		d.curr.prev.next = d.curr.next
	} else {
		d.head = d.curr.next
	}

	if d.curr.next != nil {
		d.curr.next.prev = d.curr.prev
	} else {
		d.tail = d.curr.prev
	}

	d.len--
	return nil
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	if d.curr == nil {
		return -1, fmt.Errorf("no current node")
	}
	index := 0
	node := d.head
	for node != d.curr {
		if node == nil {
			return -1, fmt.Errorf("current node not found in list")
		}
		node = node.next
		index++
	}
	return index, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	if d.len == 0 {
		return nil
	}
	node := d.tail
	if d.tail.prev != nil {
		d.tail = d.tail.prev
		d.tail.next = nil
	} else {
		d.head = nil
		d.tail = nil
	}
	d.len--
	return node
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	if d.len == 0 {
		return nil
	}
	node := d.head
	if d.head.next != nil {
		d.head = d.head.next
		d.head.prev = nil
	} else {
		d.head = nil
		d.tail = nil
	}
	d.len--
	return node
}

func (d *DoubleLinkedList) Push(c Commit) error {
	newNode := &Node{data: &c}

	if d.len == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	}

	d.len++
	return nil
}

func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n < 0 || n >= d.Len() {
		return nil, fmt.Errorf("index out of range")
	}
	var index int
	findNode := d.head
	for findNode != nil {
		if index == n {
			return findNode, nil
		}
		index++
		findNode = findNode.next
	}
	return nil, fmt.Errorf("element not found")
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	node := d.head
	for node != nil {
		if node.data.UUID == uuID {
			return node
		}
		node = node.next
	}
	return nil
}

// Search поиск коммита по message
func (d *DoubleLinkedList) Search(message string) *Node {
	node := d.head
	for node != nil {
		if node.data.Message == message {
			return node
		}
		node = node.next
	}
	return nil
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	var newDLList DoubleLinkedList
	node := d.tail
	for node != nil {
		err := newDLList.Push(*node.data)
		if err != nil {
			log.Fatal("error in Push method:", err)
		}
		node = node.prev
	}
	return &newDLList
}

type DoubleLinkedList struct {
	head *Node // начальный элемент в списке
	tail *Node // последний элемент в списке
	curr *Node // текущий элемент меняется при использовании методов next, prev
	len  int   // количество элементов в списке
}

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type Commit struct {
	Message string    `json:"message"`
	UUID    string    `json:"uuid"`
	Date    time.Time `json:"date"`
}

func GenerateData() []Commit {
	gofakeit.Seed(time.Now().UnixNano())
	n := 30
	commits := make([]Commit, n)
	for i := range commits {
		commits[i] = Commit{
			Message: gofakeit.Sentence(4),
			UUID:    gofakeit.UUID(),
			Date:    gofakeit.Date(),
		}
	}
	return commits
}

func main() {
	_, exe, _, _ := runtime.Caller(0)
	dir := filepath.Dir(exe)
	fmt.Println(dir + "/commits.json")

	var newDLList DoubleLinkedList
	err := newDLList.LoadData(dir + "/commits.json")
	if err != nil {
		log.Fatal("LoadData error:", err)
	}
	fmt.Println(newDLList.head.data.Message)
	fmt.Println(newDLList.tail.data.Message)
}
