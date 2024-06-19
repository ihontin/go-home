package main

import "fmt"

var CalculateCircleArea func(radius float64) float64 = ircleArea
var CalculateRectangleArea func(width, height float64) float64 = ctangleArea
var CalculateTriangleArea func(base, height float64) float64 = riangleArea

func ircleArea(radius float64) float64 {
	return (radius * radius) * 3.141592653589793
}
func ctangleArea(width, height float64) float64 {
	return width * height
}
func riangleArea(base, height float64) float64 {
	return (base * height) / 2
}
func main() {
	var rad, widRec, heiRec, basTri, geiTri float64
	rad = 10.0
	widRec = 7.0
	heiRec = 7.0
	basTri = 7.0
	geiTri = 7.0
	fmt.Printf("%v\n", CalculateCircleArea(rad))
	fmt.Printf("%v\n", CalculateRectangleArea(widRec, heiRec))
	fmt.Printf("%v\n", CalculateTriangleArea(basTri, geiTri))

}
