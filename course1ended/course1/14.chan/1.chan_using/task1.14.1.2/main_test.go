package main

import "testing"

//func generateData(n int) chan int {
//	ch := make(chan int)
//	go func() {
//		for i := 0; i < n; i++ {
//			ch <- i
//		}
//		close(ch)
//	}()
//	return ch
//}

func TestGenerateData(t *testing.T) {
	check := generateData(5)
	expected := []int{0, 1, 2, 3, 4}
	for _, val := range expected {
		got := <-check
		if val != got {
			t.Errorf("expected = %d, got = %d", val, got)
		}
	}
	_, ok := <-check
	if ok {
		t.Error("Expected channel to be closed")
	}
}
