package main

import (
	"fmt"
	"sync"
)

//Пример 1
//Оптимизация с использованием map для доступа за O(1).

// Создание и использование map
//m := make(map[string]int)
//m["key1"] = 1
//m["key2"] = 2

// Доступ к элементу по ключу
//value := m["key1"]

//Пример 2
//Оптимизация с использованием sync.Map для параллельного доступа.

// Создание и использование sync.Map
func main() {
	var m sync.Map
	m.Store("key1", 1)
	m.Store("key2", 2)

	// Получение значения по ключу
	value, _ := m.Load("key1")
	fmt.Println(value)
}

//Пример 3
//Оптимизация памяти с использованием sync.Pool.

// Создание и использование sync.Pool
//var pool = sync.Pool{
//	New: func() interface{} {
//		return make([]byte, 1024)
//	},
//}

// Получение объекта из пула
//buffer := pool.Get().([]byte)

// Возвращение объекта в пул
//pool.Put(buffer)

//Пример 4
//Оптимизация памяти с использованием буфера.

//import "bytes"

// Создание и использование буфера
//var buffer bytes.Buffer
//buffer.WriteString("Hello, ")
//buffer.WriteString("World!")

// Получение результирующей строки
//result := buffer.String()

//Пример 5
//Использование int вместо string.

// Использование int вместо string
//var num int = 42

// Сравнение чисел
//if num == 42 {
// Выполнение действий
//}

//Пример 6
//Использование int для флагов, которые могут быть проверены с помощью побитовых операций.

// Использование int для флагов
//const (
//	Flag1 = 1 << iota // 1
//	Flag2             // 2
//	Flag3             // 2
//	Flag4             // 4
//)
//
//func main() {
//	var flags int = 1 | 1
//	var f1 = flags & Flag2
//	var f2 = flags & Flag1
//	fmt.Println(flags, f1, f2)
//}

// Проверка флагов
//var flags int = Flag1 | Flag2

//if flags&Flag1 != 0 {
// Флаг Flag1 установлен
//}

//if flags&Flag2 != 0 {
// Флаг Flag2 установлен
//}

//Это некоторые примеры использования оптимизации функционала в Go.
//Ты можешь применять эти методы для улучшения производительности и эффективности кода.
