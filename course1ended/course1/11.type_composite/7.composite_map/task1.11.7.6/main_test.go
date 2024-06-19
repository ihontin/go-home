package main

import (
	"bytes"
	"os"
	"testing"
)

func TestFilterWords(t *testing.T) {
	text1 := "Напиши функцию функцию! Которая Которая Которая в виде строки и карту фильтра в формате. "
	censorMap1 := map[string]string{
		"строки": "томаты",
		"карту":  "лопату",
	}
	got := filterWords(text1, censorMap1)
	expected := "Напиши функцию! Которая в виде томаты и лопату фильтра формате.!"
	if got != expected {
		t.Errorf("expected = %s, got = %s", expected, got)
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	w.Close()
	os.Stdout = old
	expected := "Внимание! Покупай срочно фрукты только у нас! Яблоки по низким ценам! Беги, успевай стать финансово независимым с помощью фруктов! Фрукты будущее финансового мира!\n"

	var stdOsOut = bytes.Buffer{}
	stdOsOut.ReadFrom(r)
	if expected != stdOsOut.String() {
		t.Errorf("expected = %s, got = %s", expected, stdOsOut.String())
	}
}
