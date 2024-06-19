package main

import (
	"bytes"
	"os"
	"testing"
)

type GenText struct {
	filter             map[string]bool
	received, expected string
}

func TestFilterSentence(t *testing.T) {
	filMap := map[string]bool{"bar": true}
	genCheckT := []GenText{
		{filMap, "bar bar bar foo foo baz", "foo foo baz"},
		{filMap, "bar foo baz", "foo baz"},
		{filMap, "", ""},
		{filMap, "e e e e y y y y e e e e", "e e e e y y y y e e e e"},
	}
	for _, tex := range genCheckT {
		if got := filterSentence(tex.received, tex.filter); got != tex.expected {
			t.Errorf("expected = %s, got = %s", tex.expected, got)
		}
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	w.Close()
	os.Stdout = old
	expected := "Lorem dolor sit amet consectetur adipiscing\n"

	var stdOsOut = bytes.Buffer{}
	stdOsOut.ReadFrom(r)
	if expected != stdOsOut.String() {
		t.Errorf("expected = %s, got = %s", expected, stdOsOut.String())
	}
}
