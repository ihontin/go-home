package main

import (
	"fmt"
	"testing"
)

//go test -bench=. -benchmem -benchtime=1000x

func BenchmarkHashMap_Get(b *testing.B) {
	hTest := NewHashMap()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		hTest.Set(key, value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			key := fmt.Sprintf("key%d", i)
			_, _ = hTest.Get(key)
		}
	}
}

func BenchmarkHashMapList_Get(b *testing.B) {
	hTest := NewHashMapList()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		hTest.Set(key, value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			key := fmt.Sprintf("key%d", i)
			_, _ = hTest.Get(key)
		}
	}
}
func BenchmarkHashMapListProxy_Get(b *testing.B) {
	hTest := *NewHashMapList()
	proxyListMap := HashMapListProxy{
		realHash: hTest,
		cache:    make(map[int]interface{})}
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		hTest.Set(key, value)
	}
	hTest.Set("keyTest", "valueTest")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			_, _ = proxyListMap.Get("keyTest")
			//key := fmt.Sprintf("key%d", i)
			//_, _ = proxyListMap.Get(key)
		}
	}
}

func BenchmarkSyncMap_Get(b *testing.B) {
	hTest := NewHSyncMap()
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		hTest.Set(key, value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			key := fmt.Sprintf("key%d", i)
			//go hTest.Get(key)
			_, _ = hTest.Get(key)
		}
	}
}
