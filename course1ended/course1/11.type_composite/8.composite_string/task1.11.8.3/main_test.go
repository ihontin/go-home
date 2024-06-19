package main

import (
	"reflect"
	"testing"
)

type CountStr struct {
	received string
	len      []rune
}

func genTestRunes() []CountStr {
	var countByte = []CountStr{
		{"Привет, мир!", []rune{1055, 1088, 1080, 1074, 1077, 1090, 44, 32, 1084, 1080, 1088, 33}},
		{"Xnj nfrjt yjdsq ujl", []rune{88, 110, 106, 32, 110, 102, 114, 106, 116, 32, 121, 106, 100, 115, 113, 32, 117, 106, 108}},
		{"Что такое новый год", []rune{1063, 1090, 1086, 32, 1090, 1072, 1082, 1086, 1077, 32, 1085, 1086, 1074, 1099, 1081, 32, 1075, 1086, 1076}},
		{"", []rune{}},
		{"*", []rune{42}},
	}
	return countByte
}

func TestGetRunes(t *testing.T) {
	var getExpect = genTestRunes()
	for _, tes := range getExpect {
		got := getRunes(tes.received)
		if !reflect.DeepEqual(got, tes.len) {
			t.Errorf("expected len = %v, got len = %v", tes.len, got)
		}
	}
}

type CountStrb struct {
	received string
	len      []byte
}

func genTestByte() []CountStrb {
	var countByte = []CountStrb{
		{"Привет, мир!", []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 44, 32, 208, 188, 208, 184, 209, 128, 33}},
		{"Xnj nfrjt yjdsq ujl", []byte{88, 110, 106, 32, 110, 102, 114, 106, 116, 32, 121, 106, 100, 115, 113, 32, 117, 106, 108}},
		{"Что такое новый год", []byte{208, 167, 209, 130, 208, 190, 32, 209, 130, 208, 176, 208, 186, 208, 190, 208, 181, 32, 208, 189, 208, 190, 208, 178, 209, 139, 208, 185, 32, 208, 179, 208, 190, 208, 180}},
		{"", []byte{}},
		{"*", []byte{42}},
	}
	return countByte
}
func TestGetBytes(t *testing.T) {
	var getExpect = genTestByte()
	for _, tes := range getExpect {
		got := getBytes(tes.received)
		if !reflect.DeepEqual(got, tes.len) {
			t.Errorf("expected len = %v, got len = %v", tes.len, got)
		}
	}
}
