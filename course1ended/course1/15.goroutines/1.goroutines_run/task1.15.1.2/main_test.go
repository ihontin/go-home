package main

import (
	"fmt"
	"testing"
	"time"
)

func TestNotifyEvery(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	data := NotifyEvery(ticker, 5100*time.Millisecond, "Окончание проверки через")
	var got int
	for v := range data {
		got++
		fmt.Println(v, 5-got, "секунд")
	}
	if got != 5 {
		t.Errorf("expected = %d, got = %d", 5, got)
	}
}
