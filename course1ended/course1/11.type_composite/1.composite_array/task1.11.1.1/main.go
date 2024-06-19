package main

func sum(xs [8]int) int {
	var numSum int
	for _, n := range xs {
		numSum += n
	}
	return numSum
}
func average(xs [8]int) float64 {
	var numSum int
	for _, n := range xs {
		numSum += n
	}
	return float64(numSum) / float64(len(xs))
}
func averageFloat(xs [8]float64) float64 {
	var numSum float64
	for _, n := range xs {
		numSum += n
	}
	return numSum / float64(len(xs))
}
func reverse(xs [8]int) [8]int {
	lenNum := len(xs)
	for i, n := range xs {
		xs[lenNum-i-1] = n
	}
	return xs
}
