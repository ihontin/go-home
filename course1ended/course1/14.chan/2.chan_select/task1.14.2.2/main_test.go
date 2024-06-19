package main

import (
	"bytes"
	"os"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	timeoutFunc := timeout(3 * time.Second)
	since := time.NewTimer(3050 * time.Millisecond)
	for {
		select {
		case <-since.C:
			t.Error("function timeout did not have time to work")
			return
		default:
			if timeoutFunc() {
				return
			}
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

	expected := "Функция выполнена вовремя\n"
	var stdOut = bytes.Buffer{}
	stdOut.ReadFrom(r)
	got := stdOut.String()
	if got != expected {
		t.Errorf("expected = %s, got = %s", expected, got)
	}
}
