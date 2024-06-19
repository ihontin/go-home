package main

import (
	"container/list"
	"fmt"
	"hash/crc32"
	"sync"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type OwnMap struct {
	Key   int
	Value interface{}
}

type HashMap struct {
	hashFunk func(n string) int
	mapList  []OwnMap
}
type Option func(*HashMap)

func WithHashCRC32() Option {
	return func(h *HashMap) {
		h.hashFunk = func(n string) int {
			data := []byte(n)
			crc := crc32.ChecksumIEEE(data)
			return int(crc)
		}
	}
}
func NewHashMap(options ...Option) *HashMap {
	newMap := &HashMap{hashFunk: func(n string) int {
		data := []byte(n)
		crc := crc32.ChecksumIEEE(data)
		return int(crc)
	}}
	for _, option := range options {
		option(newMap)
	}
	return newMap
}

func (h *HashMap) Set(key string, value interface{}) {
	hash := h.hashFunk(key)
	for i, val := range h.mapList {
		if val.Key == hash {
			h.mapList[i].Value = value
			return
		}
	}
	h.mapList = append(h.mapList, OwnMap{hash, value})
}
func (h *HashMap) Get(key string) (interface{}, bool) {
	hash := h.hashFunk(key)
	for _, val := range h.mapList {
		if val.Key == hash {
			return val.Value, true
		}
	}
	return nil, false
}

//----------------------------------list

type OwnMapList struct {
	Key   int
	Value interface{}
}

type HashMapList struct {
	hashFunk func(n string) int
	dataList *list.List
}
type OptionL func(*HashMapList)

func WithHashLCRC32() OptionL {
	return func(h *HashMapList) {
		h.hashFunk = func(n string) int {
			data := []byte(n)
			crc := crc32.ChecksumIEEE(data)
			return int(crc)
		}
	}
}

func NewHashMapList(options ...OptionL) *HashMapList {
	newMap := &HashMapList{hashFunk: func(n string) int {
		data := []byte(n)
		crc := crc32.ChecksumIEEE(data)
		return int(crc)
	},
		dataList: list.New()}
	for _, option := range options {
		option(newMap)
	}
	return newMap
}

func (h *HashMapList) Set(key string, value interface{}) {
	hash := h.hashFunk(key)
	for e := h.dataList.Front(); e != nil; e = e.Next() {
		if tt, ok := e.Value.(OwnMapList); ok && tt.Key == hash {
			h.dataList.Remove(e)
			break
		}
	}

	h.dataList.PushBack(OwnMapList{hash, value})
}

func (h *HashMapList) Get(key string) (interface{}, bool) {
	hash := h.hashFunk(key)
	for e := h.dataList.Front(); e != nil; e = e.Next() {
		if tt, ok := e.Value.(OwnMapList); ok && tt.Key == hash {
			return tt.Value, true
		}
	}
	return nil, false
}

//-------------------------------------Proxy

type HashMapListProxy struct {
	realHash HashMapList
	cache    map[int]interface{}
}

func (h *HashMapListProxy) Get(key string) (interface{}, bool) {
	hash := h.realHash.hashFunk(key)
	if val, ok := h.cache[hash]; ok {
		return val, ok
	}
	if val, ok := h.realHash.Get(key); ok {
		h.cache[hash] = val
		return val, ok
	}
	return nil, false
}

//------------------------------sync.Map

type SyncMap struct {
	hashFunk func(n string) int
	sMap     *sync.Map
}

type OptionS func(*SyncMap)

func WithHashSCRC32() OptionS {
	return func(h *SyncMap) {
		h.hashFunk = func(n string) int {
			data := []byte(n)
			crc := crc32.ChecksumIEEE(data)
			return int(crc)
		}
	}
}

func NewHSyncMap(options ...OptionS) *SyncMap {
	newMap := &SyncMap{hashFunk: func(n string) int {
		data := []byte(n)
		crc := crc32.ChecksumIEEE(data)
		return int(crc)
	},
		sMap: &sync.Map{},
	}
	for _, option := range options {
		option(newMap)
	}
	return newMap
}
func (h *SyncMap) Set(key string, value interface{}) {
	hash := h.hashFunk(key)
	h.sMap.Store(hash, value)
}

func (h *SyncMap) Get(key string) (interface{}, bool) {
	hash := h.hashFunk(key)
	return h.sMap.Load(hash)
}

func main() {
	m := NewHashMap()

	m.Set("key1", "value1")
	m.Set("key2", "value2")

	if value, ok := m.Get("key1"); ok {
		fmt.Println("Key1:", value)
	} else {
		fmt.Println("Key1 not found")
	}

	if value, ok := m.Get("key2"); ok {
		fmt.Println("Key2:", value)
	} else {
		fmt.Println("Key2 not found")
	}

	if value, ok := m.Get("key3"); ok {
		fmt.Println("Key3:", value)
	} else {
		fmt.Println("Key3 not found")
	}
	listHash := *NewHashMapList()
	proxyListMap := HashMapListProxy{
		realHash: listHash,
		cache:    make(map[int]interface{})}
	listHash.Set("key1", "value1")
	listHash.Set("key2", "value2")

	if value, ok := proxyListMap.Get("key1"); ok {
		fmt.Println("Key1:", value)
	} else {
		fmt.Println("Key1 not found")
	}

	if value, ok := proxyListMap.Get("key2"); ok {
		fmt.Println("Key2:", value)
	} else {
		fmt.Println("Key2 not found")
	}
	if value, ok := proxyListMap.Get("key3"); ok {
		fmt.Println("Key3:", value)
	} else {
		fmt.Println("Key3 not found")
	}
	if value, ok := proxyListMap.Get("key2"); ok {
		fmt.Println("Key2:", value)
	} else {
		fmt.Println("Key2 not found")
	}

	syncTest := NewHSyncMap()
	fmt.Println("----------------------------------")
	syncTest.Set("key1", "value1")
	syncTest.Set("key2", "value2")

	if value, ok := syncTest.Get("key1"); ok {
		fmt.Println("Key1:", value)
	} else {
		fmt.Println("Key1 not found")
	}

	if value, ok := syncTest.Get("key2"); ok {
		fmt.Println("Key2:", value)
	} else {
		fmt.Println("Key2 not found")
	}
	if value, ok := syncTest.Get("key3"); ok {
		fmt.Println("Key3:", value)
	} else {
		fmt.Println("Key3 not found")
	}
}
