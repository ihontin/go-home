package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

// https://github.com/ptflp/kata_tasks
// 1
func tribFun(n int) int {
	if n <= 1 {
		return n
	}
	a, b, res := 0, 1, 1
	for i := 1; i <= n-2; i++ {
		a, b, res = b, res, a+b+res
	}
	return res
}

// 2
func conc(ar []int) []int {
	arr := make([]int, 0, len(ar)*2)
	arr = append(ar, ar...)
	return arr
}

// 3
func convertTemperature(celsius float64) []float64 {
	relvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00
	return []float64{relvin, fahrenheit}
}

// 4
func buildArray(nums []int) []int {
	var newNums = make([]int, 0, len(nums))
	for _, val := range nums {
		newNums = append(newNums, nums[val])
	}
	return newNums
}

// 5
func numberOfMatches(n int) int {
	if n < 2 {
		return 0
	}
	if n%2 == 0 {
		return (n / 2) + numberOfMatches(n/2)
	}
	return (n / 2) + numberOfMatches((n/2)+1)
}

// 6
func findIndex(l string) int {
	serchL := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for i := range serchL {
		if l == serchL[i] {
			return i
		}
	}
	return -1
}
func uniqueMorseRepresentations(words []string) int {
	morzWords := make(map[string]bool, len(words))
	morze := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, word := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			var mWord string
			for _, lett := range w {
				mWord += morze[findIndex(string(lett))]
			}
			mutex.Lock()
			morzWords[mWord] = true
			mutex.Unlock()
		}(word)
	}
	wg.Wait()
	return len(morzWords)
}

// 7
func defangIPaddr(address string) string {
	r := regexp.MustCompile(`\.`)
	outS := r.ReplaceAllString(address, "[.]")
	return outS
}

// 8
func findKthPositive(arr []int, k int) int {
	arrIndex := 0
	count := 0
	for i := 1; k >= count; i++ {
		if i != arr[arrIndex] {
			count++
		} else if i == arr[arrIndex] && arrIndex < len(arr)-1 {
			arrIndex++
		}
		if k == count {
			return i
		}
	}
	return -1
}

// 9
func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, val := range operations {
		lenStr := utf8.RuneCountInString(val)
		if string(val[0]) == "-" || string(val[lenStr-1]) == "-" {
			x--
			continue
		}
		x++
	}
	return x
}

// 10
func shuffle(nums []int, n int) []int {
	newList := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		newList = append(newList, nums[i], nums[n+i])
	}
	return newList
}

// 11
func runningSum(nums []int) []int {
	curSum := 0
	for i, val := range nums {
		curSum += val
		nums[i] = curSum
	}
	return nums
}

// 12
func numIdenticalPairs(nums []int) int {
	count := 0
	for i, num := range nums {
		for _, n := range nums[i+1:] {
			if num == n {
				count++
			}
		}
		if i == len(nums)-2 {
			break
		}
	}
	return count
}

// 13. Драгоценности и камни Input: jewels = "aA", stones = "aAAbbbb"
func numJewelsInStones(jewels string, stones string) int {
	coutn := 0
	for _, stone := range stones {
		for _, jewel := range jewels {
			if stone == jewel {
				coutn++
				continue
			}
		}
	}
	return coutn
}

// 14. Богатство самого богатого клиента
func maximumWealth(accounts [][]int) int {
	traj := 0
	for _, val := range accounts {
		count := 0
		for _, v := range val {
			count += v
		}
		if count > traj {
			traj = count
		}
	}
	return traj
}

// 15. Дизайн парковочной системы
type ParkingSystem struct {
	big, medium, small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big--
			return true
		} else {
			return false
		}
	case 2:
		if this.medium > 0 {
			this.medium--
			return true
		} else {
			return false
		}
	case 3:
		if this.small > 0 {
			this.small--
			return true
		} else {
			return false
		}
	default:
		return false
	}
}

// 16. Наименьшее четное кратное
func smallestEvenMultiple(n int) int {
	for i := n; ; i += n {
		if i%2 == 0 {
			return i
		}
	}
}

// 17. Максимальное количество найденных слов в предложениях
func mostWordsFound(sentences []string) int {
	sum := 0
	for _, val := range sentences {
		s := strings.Count(val, " ") + 1
		if s > sum {
			sum = s
		}
	}
	return sum
}

// 18. Разница между суммой элементов и суммой цифр массива
func differenceOfSum(nums []int) int {
	sumAll := 0
	sumUnick := 0
	for _, val := range nums {
		sumAll += val
		for val != 0 {
			sumUnick += val % 10
			val = val / 10
		}
	}
	return sumAll - sumUnick
}

// 19
func minimumSum(num int) int {
	var sum int
	var unit, tens, hundr, thous string
	val := num
	min := num%100 + num/100
	var fourL = make([]string, 0, 4)
	unit, val = strconv.Itoa(val%10), val/10
	tens, val = strconv.Itoa(val%10), val/10
	hundr, val = strconv.Itoa(val%10), val/10
	thous = strconv.Itoa(val % 10)
	fourL = append(fourL, unit, tens, hundr, thous)
	for i := 0; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for q := 0; q < 2; q++ {
				first, _ := strconv.Atoi(fourL[0])
				first2, _ := strconv.Atoi(fourL[0] + fourL[j])
				var sec, sec2 int
				switch j {
				case 1:
					sec, _ = strconv.Atoi(fourL[j] + fourL[2] + fourL[3])
					sec2, _ = strconv.Atoi(fourL[2] + fourL[3])
					if q > 0 {
						sec, _ = strconv.Atoi(fourL[j] + fourL[3] + fourL[2])
						sec2, _ = strconv.Atoi(fourL[3] + fourL[2])
					}
				case 2:
					sec, _ = strconv.Atoi(fourL[j] + fourL[1] + fourL[3])
					sec2, _ = strconv.Atoi(fourL[1] + fourL[3])
					if q > 0 {
						sec, _ = strconv.Atoi(fourL[j] + fourL[3] + fourL[1])
						sec2, _ = strconv.Atoi(fourL[3] + fourL[1])
					}
				case 3:
					sec, _ = strconv.Atoi(fourL[j] + fourL[1] + fourL[2])
					sec2, _ = strconv.Atoi(fourL[1] + fourL[2])
					if q > 0 {
						sec, _ = strconv.Atoi(fourL[j] + fourL[2] + fourL[1])
						sec2, _ = strconv.Atoi(fourL[2] + fourL[1])
					}
				}
				sum = first + sec
				if min > sum {
					min = sum
				}
				sum = first2 + sec2
				if min > sum {
					min = sum
				}
			}
		}
		fourL = append(fourL[1:], fourL[0])
	}
	return min
}

// 20
func kidsWithCandies(candies []int, extraCandies int) []bool {
	biger := 0
	boolList := make([]bool, 0, len(candies))
	for _, val := range candies {
		if val > biger {
			biger = val
		}
	}
	for _, val := range candies {
		if val+extraCandies >= biger {
			boolList = append(boolList, true)
			continue
		}
		boolList = append(boolList, false)
	}
	return boolList
}

func main() {
	//1
	//fmt.Println(tribFun(1))
	//2
	//fmt.Println(conc([]int{1, 2, 3, 4, 5}))
	//3
	//fmt.Println(convertTemperature(36.50))
	//4
	//fmt.Println(buildArray([]int{0, 2, 1, 5, 3, 4}))
	//5
	//fmt.Println(numberOfMatches(5))
	//6
	//fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
	//7
	//fmt.Println(defangIPaddr("ad.dre.ss"))
	//8
	//fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
	//9
	//fmt.Println(finalValueAfterOperations([]string{"++X", "++X", "X ++"}))
	//10
	//fmt.Println(shuffle([]int{1, 2, 3, 4, 4, 3, 2, 1}, 4))
	//11
	//fmt.Println(runningSum([]int{3, 4, 6, 16, 17}))
	//12
	//fmt.Println(numIdenticalPairs([]int{1, 2, 3}))
	//13
	//fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	//14
	//fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 1, 2}}))
	//15
	//16
	//fmt.Println(smallestEvenMultiple(8))
	//17
	//fmt.Println(mostWordsFound([]string{"please wait", "continue to fight", "continue to win"}))
	//18
	//fmt.Println(differenceOfSum([]int{1, 2, 3, 4}))
	//19
	//fmt.Println(minimumSum(7890))
	//20
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
