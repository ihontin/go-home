package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} {
		return &Person{}
	},
} // sync.Pool of Person

func main() {
	newPerson := &Person{12}
	personPool.Put(newPerson)

	fromPool := personPool.Get().(*Person)
	fmt.Println(fromPool.Age)
}
