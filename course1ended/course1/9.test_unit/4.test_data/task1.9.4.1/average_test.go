package main

import (
	"math/rand"
	"testing"
	"time"
)

func TestAverage(t *testing.T) {
	for i := 100; i > 0; i-- {
		randList := generator()
		var randSum float64
		for _, num := range randList {
			randSum += num
		}
		got := float64(randSum / 1000)
		if expected := average(randList); expected != got {
			t.Errorf("expected = %f, got = %f", expected, got)
		}
	}
}

func generator() []float64 {
	rand.Seed(time.Now().UnixNano())
	var listNums = make([]float64, 0, 1000)
	for i := 0; i < 1000; i++ {
		listNums = append(listNums, float64(rand.Intn(999)))
	}
	return listNums
}
