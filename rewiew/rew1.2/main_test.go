package main

import (
	"math/rand"
	"testing"
	"time"
)

type FizBuz struct {
	received int
	expected string
}

func fizzBuzzTest(x int) string {
	switch {
	case x%3 == 0 && x%5 == 0:
		return "fizz buzz"
	case x%3 != 0 && x%5 == 0:
		return "buzz"
	case x%3 == 0 && x%5 != 0:
		return "fizz"
	default:
		return "wrong number"
	}
}
func randGen() []FizBuz {
	rand.Seed(time.Now().UnixNano())
	oneList := FizBuz{}
	var FizBuzListStruct = []FizBuz{}
	for i := 0; i < 1000; i++ {
		oneList.received = rand.Intn(1000)
		oneList.expected = fizzBuzzTest(oneList.received)
		FizBuzListStruct = append(FizBuzListStruct, oneList)
	}
	return FizBuzListStruct
}
func TestFizzBuzz(t *testing.T) {
	fizzBuzzTests := randGen()

	for _, val := range fizzBuzzTests {
		if got := fizzBuzz(val.received); got != val.expected {
			t.Errorf("received = %d, expected = %s, got = %s", val.received, val.expected, got)
		}
	}
}
