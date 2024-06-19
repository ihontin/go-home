package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	fmt.Println(Max(1, 5))
	var a, b uint
	a, b = 4, 7
	fmt.Println(Equal(a, b))
	fmt.Println(MaxOrderP(a, b))
	fmt.Println(MaxOrder("a", "b"))
	var lookInList = []string{",", ".", "!"}
	fmt.Println(findInList(".", lookInList))
	var maxInList = []float64{5.6, 7.03, 7.06}
	fmt.Println(findMax(maxInList))
	sum := func(x, y float64) float64 { return x + y } // create a function that returns the sum of two values
	mul := func(x, y float64) float64 { return x * y } // create a function that returns the multiple of two values
	res := Reduce(maxInList, sum, 0)                   //
	fmt.Println(res, Reduce(maxInList, mul, 1))
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Equal constraint comparable позволяет сравнить два числа используя ==
func Equal[T comparable](x, y T) bool {
	if x == y {
		return true
	}
	return false
}

// OrderedP create private constraint
type OrderedP interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// MaxOrderP use private constraint
func MaxOrderP[O OrderedP](x, y O) O {
	if x < y {
		return y
	}
	return x
}

// MaxOrder golang.org/x/exp/constraints
func MaxOrder[O constraints.Ordered](x, y O) O {
	if x < y {
		return y
	}
	return x
}

// позволяет сравнивать любые типы данных
func findInList[T comparable](n T, l []T) bool {
	for _, val := range l {
		if val == n {
			return true
		}
	}
	return false
}

// OrderedIntFloat create private constraint
type OrderedIntFloat interface {
	constraints.Integer | constraints.Float
}

func findMax[T OrderedIntFloat](l []T) T {
	var max = l[0]
	for _, val := range l {
		if val > max {
			max = val
		}
	}
	return max
}

func Reduce[T any](list []T, accumulator func(T, T) T, init T) T {
	for _, el := range list {
		init = accumulator(init, el)
	}
	return init
}
