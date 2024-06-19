package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGetScanner(t *testing.T) {
	buffer := bytes.NewBufferString("Hello, World!")
	expected := "*bufio.Scanner"
	if got := fmt.Sprintf("%v", getScanner(buffer)); got != expected {
		fmt.Errorf("expected %s, but got %s", expected, got)
	}
}
