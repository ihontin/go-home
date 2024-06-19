package main

import (
	"fmt"
	"sort"
	"strings"
)

// 21
func subtractProductAndSum(n int) int {
	var sum int
	var mult = 1
	for i := 0; i < 6; i++ {
		if n < 1 {
			break
		}
		num := n % 10
		sum += num
		mult *= num
		n = n / 10
	}
	return mult - sum
}

// 22
func smallerNumbersThanCurrent(nums []int) []int {
	var lesList = make([]int, 0, len(nums))
	for i, val := range nums {
		count := 0
		for j, eq := range nums {
			if i == j {
				continue
			}
			if val > eq {
				count++
			}
		}
		lesList = append(lesList, count)
	}
	return lesList
}

// 23
func interpret(command string) string {
	command = strings.ReplaceAll(command, "()", "o")
	command = strings.ReplaceAll(command, "(al)", "al")
	return command
}

// 24
func decode(encoded []int, first int) []int {
	var res = make([]int, len(encoded)+1)
	res[0] = first
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] ^ encoded[i-1]
	}
	return res
}

// 25
func createTargetArray(nums []int, index []int) []int {
	res := make([]int, 0, len(nums))
	for a, i := range index {
		res = append(res[:i], append([]int{nums[a]}, res[i:]...)...)
		fmt.Println(a, res)
	}
	return res
}

// 26
func decompressRLElist(nums []int) []int {
	outL := make([]int, 0, 100)
	for i := 1; i < len(nums); i += 2 {
		for j := 0; j < nums[i-1]; j++ {
			outL = append(outL, nums[i])
		}
	}
	return outL
}

// 27 R82 L76
func balancedStringSplit(s string) int {
	var count int
	var r, l int
	for _, lr := range s {
		if lr == 82 {
			r++
		} else if lr == 76 {
			l++
		}
		if r == l {
			count++
		}
	}
	return count
}

// 28
func countDigits(num int) int {
	var count int
	copyNum := num
	for num > 0 {
		divNum := num % 10
		num /= 10
		if copyNum%divNum == 0 {
			count++
		}
	}
	return count
}

// 29
func xorOperation(n int, start int) int {
	var outL int
	for i := 0; i < n; i++ {
		if i == 0 {
			outL = start + 2*i
			continue
		}
		outL ^= start + 2*i
	}
	return outL
}

// 30
func countGoodTriplets(arr []int, a int, b int, c int) int {
	var count int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for q := j + 1; q < len(arr); q++ {
				if abs(arr[j]-arr[q]) <= b && abs(arr[i]-arr[q]) <= c {
					count++
				}
			}
		}
	}
	return count
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 31
func sortPeople(names []string, heights []int) []string {
	var sortMap = make(map[int]string, len(names))
	var sortList = make([]string, len(names))
	for i := range names {
		sortMap[heights[i]] = names[i]
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] < heights[j]
	})
	for i := len(heights) - 1; i >= 0; i-- {
		sortList[len(heights)-1-i] = sortMap[heights[i]]
	}
	return sortList
}

func main() {
	//21
	//fmt.Println(subtractProductAndSum(4421))
	//22
	//fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	//23
	//fmt.Println(interpret("G()()()()(al)"))
	//24
	//fmt.Println(decode([]int{6, 2, 7, 3}, 4))
	//25
	//fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}))
	//26
	//fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
	//27
	//fmt.Println(balancedStringSplit("LLLLRRRR"))
	//28
	//fmt.Println(countDigits(121))
	//29
	//fmt.Println(xorOperation(4, 3))
	//30
	//fmt.Println(countGoodTriplets([]int{3,0,1,1,9,7}, 7, 2, 3))
	//31
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170}))

}
