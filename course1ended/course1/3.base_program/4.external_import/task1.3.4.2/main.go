package main

import (
	"fmt"
	"github.com/ksrof/gocolors"
	costum "github.com/mewzax/gocolors"
)

// ColorizeRed — функция, которая принимает строку и возвращает ее в красном цвете.
func ColorizeRed(a string) string {
	return gocolors.Red(a, "")
}

// ColorizeGreen — функция, которая принимает строку и возвращает ее в зеленом цвете.
func ColorizeGreen(a string) string {
	return gocolors.Green(a, "")
}

// ColorizeBlue — функция, которая принимает строку и возвращает ее в синем цвете.
func ColorizeBlue(a string) string {
	return gocolors.Blue(a, "")
}

// ColorizeYellow — функция, которая принимает строку и возвращает ее в желтом цвете.
func ColorizeYellow(a string) string {
	return gocolors.Yellow(a, "")
}

// ColorizeMagenta — функция, которая принимает строку и возвращает ее в пурпурном цвете.
func ColorizeMagenta(a string) string {
	return gocolors.Magenta(a, "")
}

// ColorizeCyan — функция, которая принимает строку и возвращает ее в голубом цвете.
func ColorizeCyan(a string) string {
	return gocolors.Cyan(a, "")
}

// ColorizeWhite — функция, которая принимает строку и возвращает ее в белом цвете.
func ColorizeWhite(a string) string {
	return gocolors.White(a, "")
}

// ColorizeCustom — функция, которая принимает строку и RGB значения цвета и возвращает ее в пользовательском цвете.
func ColorizeCustom(a string, r, g, b uint8) string {
	return costum.Colorize(costum.RGB(int(r), int(g), int(b)), a)
}

func main() {
	var inColor string
	var r, g, b uint8

	fmt.Scanln(&inColor)
	fmt.Scanln(&r)
	fmt.Scanln(&g)
	fmt.Scanln(&b)
	fmt.Println(ColorizeRed(inColor))
	fmt.Println(ColorizeGreen(inColor))
	fmt.Println(ColorizeBlue(inColor))
	fmt.Println(ColorizeYellow(inColor))
	fmt.Println(ColorizeMagenta(inColor))
	fmt.Println(ColorizeCyan(inColor))
	fmt.Println(ColorizeWhite(inColor))
	fmt.Println(ColorizeCustom(inColor, r, g, b))
}
