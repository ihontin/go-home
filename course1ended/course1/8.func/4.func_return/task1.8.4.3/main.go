package main

import "fmt"

// CalculateStockValue (price float64, quantity int) (float64, float64),
// которая вычисляет общую стоимость акций, умножая цену на количество, и возвращает два значения:
//
//	общую стоимость акций и цену одной акции. Функция должна находиться в пакете main.
func CalculateStockValue(price float64, quantity int) (float64, float64) {
	return price * float64(quantity), price
}

func main() {
	fmt.Println(CalculateStockValue(9.3, 4))
}
