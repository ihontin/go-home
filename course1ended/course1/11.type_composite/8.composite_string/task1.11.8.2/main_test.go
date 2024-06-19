package main

import (
	"reflect"
	"testing"
)

type CountStr struct {
	received string
	len      int
}

func genTestByte() []CountStr {
	var countByte = []CountStr{
		{"Привет, мир!", 21},
		{"Xnj nfrjt yjdsq ujl", 19},
		{"Что такое новый год", 35},
		{"", 0},
		{"*", 1},
	}
	return countByte
}

func TestGetStringHeader(t *testing.T) {
	var getExpect = genTestByte()
	var header reflect.StringHeader
	for _, tes := range getExpect {
		header = getStringHeader(tes.received)
		gotL := header.Len
		if gotL != tes.len {
			t.Errorf("expected len = %v, got len = %v", tes.len, gotL)
		}
	}
}
