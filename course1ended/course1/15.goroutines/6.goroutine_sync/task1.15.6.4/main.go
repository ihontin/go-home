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

type Cache struct {
	data  map[string]*User
	mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]*User),
	}
}

func (c *Cache) Set(key string, user *User) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = user
}

func (c *Cache) Get(key string) *User {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.data[key]
}

func keyBuilder(keys ...string) string {
	return strings.Join(keys, ":")
}

func main() {
	cache := NewCache()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.Set(keyBuilder("user", strconv.Itoa(i)), &User{
				ID:   i,
				Name: fmt.Sprint("user-", i),
			})
		}(i)
	}
	//time.Sleep(1 * time.Second)
	wg.Wait()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(cache.Get(keyBuilder("user", strconv.Itoa(i))))
		}(i)
	}
	wg.Wait()
}
