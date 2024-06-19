package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID   int
	Name string
}

// Cache - структура для конкурентного доступа к данным других объектов
type Cache struct {
	mutex sync.RWMutex           // Поле mutex позволяет одновременно читать
	data  map[string]interface{} // Поле data может хранить любой объект
}

// NewCache Конструктор для создания экземпляра Cache
func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

// Set - метод добавляющий данные в Cache
func (c *Cache) Set(key string, user *User) {
	c.mutex.Lock()         // Блокирует доступ другим горутинам к следующим строкам кода
	defer c.mutex.Unlock() // Снимает блокировку после окончания работы функции
	c.data[key] = user
}

// Get - метод возвращает данные из Cache по ключу
func (c *Cache) Get(key string) interface{} {
	c.mutex.RLock()         // Количество горутин для чтения не ограничено, но при записи произойдет блокировка
	defer c.mutex.RUnlock() // Снимает блокировку после окончания работы функции
	return c.data[key]
}

// keyBuilder создает ключ для Cache в нужном формате
func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

// GetUser утверждает тип интерфейса и фозвращает значение в виде структуры
func GetUser(i interface{}) *User {
	return i.(*User)
}

func main() {
	cache := NewCache()   // создание экземпляра структуры Cache
	var wg sync.WaitGroup // создание группы ожидания
	// цикл создаст максимально возможное количество горутин
	for i := 0; i < 100; i++ {
		wg.Add(1) // добавление к группе ожидания
		//передаем индекс итерации в горутину
		go func(i int) {
			defer wg.Done() //при выходе из функции удалит одно значение из группы ожидания
			// добавляем данные Юзера в экземпляр cache
			cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
				ID:   i,
				Name: fmt.Sprint("user-", i),
			})
		}(i)
	}
	wg.Wait() // останавливает основную горутидо до опустошения группы ожидания
	// цикл создаст максимально возможное количество горутин
	for i := 0; i < 100; i++ {
		wg.Add(1) // добавление к группе ожидания
		//передаем индекс итерации в горутину
		go func(i int) {
			defer wg.Done() //при выходи из функции удалит одно значение из группы ожидания
			// возвращает данные Юзера из Cache по ключу
			raw := cache.Get(keyBuilder("user", strconv.Itoa(i)))
			fmt.Println(GetUser(raw)) // достаем значение из интерфейса
		}(i)
	}
	wg.Wait() // останавливает основную горутину до опустошения группы ожидания
}
