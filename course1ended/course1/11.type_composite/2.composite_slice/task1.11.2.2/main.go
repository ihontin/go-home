package main

func MaxDifference(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	var min, max = 1<<63 - 1, -1 << 63
	for _, number := range numbers {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
	}
	return max - min
}
