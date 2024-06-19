package main

func mutateNums(sl []int, idx int, val int) {
	sl[idx] = val
}
func main() {
	nums := make([]int, 1, 512)
	mutateNums(nums, 0, 666)
}
