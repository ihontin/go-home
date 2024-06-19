package main

import (
	"fmt"
	"testing"
)

func TestPreparePrint(t *testing.T) {
	testCheck := getUsers()
	var expected string
	for _, user := range testCheck {
		expected += fmt.Sprintf("Имя: %s, Возраст: %d\n", user.Name, user.Age)
	}
	if got := preparePrint(testCheck); got != expected {
		t.Errorf("expected = %v, got = %v", expected, got)
	}
}
