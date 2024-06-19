package main

import "fmt"

// CalculateSimpleMovingAverage (period int, data ...float64) []float64,
// которая вычисляет простое скользящее среднее для заданного периода и набора данных.
func CalculateSimpleMovingAverage(period int, data ...float64) []float64 {
	if period < 1 {

		return []float64{}
	}
	listSma := make([]float64, 0, len(data)/period)
	n := 0
	var sumDatas float64
	for _, da := range data {
		n++
		sumDatas += da
		if n == period {
			n = 0
			listSma = append(listSma, sumDatas/float64(period))
			sumDatas = 0
		}
	}
	return listSma
}

func main() {
	sma1 := CalculateSimpleMovingAverage(2, 3.0, 3.0, 3.0, 3.0, 3.0, 3.0, 0.3)
	sma2 := CalculateSimpleMovingAverage(5)
	sma3 := CalculateSimpleMovingAverage(0, 0.0)
	sma4 := CalculateSimpleMovingAverage(3, 5.0, 5.0, 5.0, 5.0, 5.0, 5.0, 5.5, 5.0)
	fmt.Println(sma1, sma2, sma3, sma4)
}
