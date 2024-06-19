package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTrySend(t *testing.T) {
	ch := make(chan int)
	go func(chans chan int) {
		fmt.Println("we are in")
		<-chans
		close(chans)
	}(ch)
	time.Sleep(time.Millisecond * 200)
	got := trySend(ch, 5)
	if !got {
		t.Errorf("expected = %t, got = %t", false, got)
	}
}
