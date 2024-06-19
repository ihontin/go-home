package main

func average(xs []float64) float64 {
	var sum float64
	for _, val := range xs {
		sum += val
	}
	return sum / float64(len(xs))
}
