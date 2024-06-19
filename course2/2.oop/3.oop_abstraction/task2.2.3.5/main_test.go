package main

import (
	"reflect"
	"testing"
	"time"
)

// go test -coverprofile=co.out
// go test -func=co.out

type KeyValTest struct {
	Key string
	Val interface{}
}

func TestWithHashCRC64(t *testing.T) {
	keyList := []KeyValTest{
		{"", ""},
		{"kay", "value"},
	}
	hashM := &HashMap{}
	WithHashCRC64()(hashM)
	hashM.mapList = make(map[int]*OwnMap)
	for _, exp := range keyList {
		hashM.Set(exp.Key, exp.Val)
		got, _ := hashM.Get(exp.Key)
		if got != exp.Val {
			t.Errorf("expected = %s, got = %s", exp.Val, got)
		}
	}
}
func TestWithHashCRC32(t *testing.T) {
	keyList := []KeyValTest{
		{"", ""},
		{"kay", "value"},
	}
	hashM := &HashMap{}
	WithHashCRC32()(hashM)
	hashM.mapList = make(map[int]*OwnMap)
	for _, exp := range keyList {
		hashM.Set(exp.Key, exp.Val)
		got, _ := hashM.Get(exp.Key)
		if got != exp.Val {
			t.Errorf("expected = %s, got = %s", exp.Val, got)
		}
	}
}
func TestWithHashCRC16(t *testing.T) {
	keyList := []KeyValTest{
		{"", ""},
		{"kay", "value"},
	}
	hashM := &HashMap{}
	WithHashCRC16()(hashM)
	hashM.mapList = make(map[int]*OwnMap)
	for _, exp := range keyList {
		hashM.Set(exp.Key, exp.Val)
		got, _ := hashM.Get(exp.Key)
		if got != exp.Val {
			t.Errorf("expected = %s, got = %s", exp.Val, got)
		}
	}
}
func TestWithHashCRC8(t *testing.T) {
	keyList := []KeyValTest{
		{"", ""},
		{"kay", "value"},
	}
	hashM := &HashMap{}
	WithHashCRC8()(hashM)
	hashM.mapList = make(map[int]*OwnMap)
	for _, exp := range keyList {
		hashM.Set(exp.Key, exp.Val)
		got, _ := hashM.Get(exp.Key)
		if got != exp.Val {
			t.Errorf("expected = %s, got = %s", exp.Val, got)
		}
	}
}

func TestNewHashMap(t *testing.T) {
	keyList := []KeyValTest{
		{"", ""},
		{"kay", "value"},
	}
	m := NewHashMap(2, WithHashCRC8())
	for _, exp := range keyList {
		m.Set(exp.Key, exp.Val)
		got, _ := m.Get(exp.Key)
		if got != exp.Val {
			t.Errorf("expected = %s, got = %s", exp.Val, got)
		}
	}
}

func TestHashMap_Set(t *testing.T) {
	keyList := []KeyValTest{
		{"", ""},
		{"kay", 8},
	}
	m := NewHashMap(2, WithHashCRC8())
	for _, exp := range keyList {
		m.Set(exp.Key, exp.Val)
		got, _ := m.Get(exp.Key)
		if got != exp.Val {
			t.Errorf("expected = %s, got = %s", exp.Val, got)
		}
	}
}

func TestHashMap_Get(t *testing.T) {
	keyList := []KeyValTest{
		{"", ""},
		{"kay", 8},
		{"rrr", "rrr"},
	}
	m := NewHashMap(2, WithHashCRC8())
	for i, exp := range keyList {
		m.Set(exp.Key, exp.Val)
		k := exp.Key
		if i == 2 {
			k = "not in list"
			got, ok := m.Get(k)
			if ok {
				t.Errorf("expected = %s, got = %s", exp.Val, got)
			}
			break
		}
		got, ok := m.Get(k)
		if got != exp.Val || !ok {
			t.Errorf("expected = %s, got = %s", exp.Val, got)
		}
	}
}

func testingFunc() {
	f := make(map[int]int, 1)
	for i := 0; i < 10000; i++ {
		f[i] = i * 2
	}
}

func TestMeassureTime(t *testing.T) {
	start := time.Now()
	f := make(map[int]int, 1)
	for i := 0; i < 10000; i++ {
		f[i] = i * 2
	}
	expected := time.Since(start)
	got := MeassureTime(testingFunc)
	if reflect.TypeOf(expected) != reflect.TypeOf(got) {
		t.Errorf("expected = %s, got = %s", expected, got)
	}
}
