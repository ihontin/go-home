package main

import (
	"fmt"
	"testing"
)

func TestPreparePrint(t *testing.T) {
	testCheck := getAnimals()
	var expected string
	for _, an := range testCheck {
		expected += fmt.Sprintf("Тип: %s, Имя: %s, Возраст: %d\n", an.Type, an.Name, an.Age)
	}
	if got := preparePrint(testCheck); got != expected {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}
