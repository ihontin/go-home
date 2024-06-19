package main

import (
	"fmt"
)

func fgor(d float64) (h int) {
	h = int(d)
	return
}

func main() {
	fmt.Println(fgor(43))
}
