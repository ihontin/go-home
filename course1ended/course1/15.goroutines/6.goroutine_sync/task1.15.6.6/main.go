package main

import "sync"

type Cache struct {
	data sync.Map
}

func (c *Cache) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	val, ok := c.data.Load(key)
	return val, ok
}
