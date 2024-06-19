package main

import "testing"

type GenText struct {
	received, expected string
}

func TestGetUniqueWords(t *testing.T) {
	genCheckT := []GenText{
		{"bar bar bar foo foo baz", "bar foo baz"},
		{"bar foo baz", "bar foo baz"},
		{"", ""},
		{"e e e e y y y y e e e e", "e y"},
	}
	for _, tex := range genCheckT {
		if got := getUniqueWords(tex.received); got != tex.expected {
			t.Errorf("expected = %s, got = %s", tex.expected, got)
		}
	}
}
