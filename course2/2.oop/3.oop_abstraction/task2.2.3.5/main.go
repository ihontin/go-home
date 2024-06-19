package main

import (
	"github.com/sigurn/crc16"
	"github.com/sigurn/crc8"
	"hash/crc32"
	"hash/crc64"
	"time"
)

type Option func(*HashMap)

func WithHashCRC64() Option {
	return func(h *HashMap) {
		h.hashFunk = func(n string) int {
			data := []byte(n)
			crcT := crc64.MakeTable(crc64.ISO)
			crc := crc64.Checksum(data, crcT)
			return int(crc)
		}
	}
}

func WithHashCRC32() Option {
	return func(h *HashMap) {
		h.hashFunk = func(n string) int {
			data := []byte(n)
			crc := crc32.ChecksumIEEE(data)
			return int(crc)
		}
	}
}

func WithHashCRC16() Option {
	return func(h *HashMap) {
		h.hashFunk = func(n string) int {
			data := []byte(n)
			table := crc16.MakeTable(crc16.CRC16_MODBUS)
			crc := crc16.Checksum(data, table)
			return int(crc)
		}
	}
}

func WithHashCRC8() Option {
	return func(h *HashMap) {
		h.hashFunk = func(n string) int {
			data := []byte(n)
			table := crc8.MakeTable(crc8.CRC8_MAXIM)
			crc := crc8.Checksum(data, table)
			return int(crc)
		}
	}
}

type OwnMap struct {
	key   string
	value interface{}
}

type HashMap struct {
	hashFunk func(n string) int
	mapList  map[int]*OwnMap
}

func NewHashMap(size int, options ...Option) *HashMap {
	outMap := &HashMap{mapList: make(map[int]*OwnMap, size)}
	for _, option := range options {
		option(outMap)
	}
	return outMap
}

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

func (h *HashMap) Set(key string, value interface{}) {
	k := h.hashFunk(key)
	h.mapList[k] = &OwnMap{
		key, value,
	}
}

func (h *HashMap) Get(key string) (interface{}, bool) {
	k := h.hashFunk(key)
	if _, ok := h.mapList[k]; !ok {
		return nil, false
	}
	return h.mapList[k].value, true
}

func MeassureTime(a func()) time.Duration {
	start := time.Now()
	a()
	return time.Since(start)
}

//func main() {
//	m := NewHashMap(16, WithHashCRC64())
//	since := MeassureTime(func() {
//
//		m.Set("key", "value")
//
//		if value, ok := m.Get("key"); ok {
//			fmt.Println(value)
//		}
//	})
//	fmt.Println(since)
//
//	m = NewHashMap(16, WithHashCRC32())
//	since = MeassureTime(func() {
//		m.Set("key", "value")
//
//		if value, ok := m.Get("key"); ok {
//			fmt.Println(value)
//		}
//	})
//	fmt.Println(since)
//
//	m = NewHashMap(16, WithHashCRC16())
//	since = MeassureTime(func() {
//		m.Set("key", "value")
//
//		if value, ok := m.Get("key"); ok {
//			fmt.Println(value)
//		}
//	})
//	fmt.Println(since)
//
//	m = NewHashMap(16, WithHashCRC8())
//	since = MeassureTime(func() {
//		m.Set("key", "value")
//
//		if value, ok := m.Get("key"); ok {
//			fmt.Println(value)
//		}
//	})
//	fmt.Println(since)
//}
