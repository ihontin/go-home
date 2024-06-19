package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// ----------------------------------------------1
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	sum, _ := sumLeaves(root, 0)
	return sum
}
func sumLeaves(node *TreeNode, depth int) (int, int) {
	if node == nil {
		return 0, 0
	}
	depth++
	if node.Left == nil && node.Right == nil {
		return node.Val, depth
	}
	lSum, lDepth := sumLeaves(node.Left, depth)
	rSum, rDepth := sumLeaves(node.Right, depth)
	if lDepth == rDepth {
		return lSum + rSum, lDepth
	} else if lDepth > rDepth {
		return lSum, lDepth
	} else {
		return rSum, rDepth
	}
}

// -------------------------------------------------2
func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		fmt.Println(score[i][k] < score[j][k])
		return score[i][k] > score[j][k]
	})
	return score
}

// ----------------------------------------- 3
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	next, newHead := head.Next, head
	for next.Next != nil {
		if next.Val == 0 {
			newHead.Next = next
			newHead = next
		}
		newHead.Val += next.Val
		next = next.Next
	}
	newHead.Next = nil
	return head
}

// ----------------------------------------------------------4
//
//	type TreeNode struct {
//		     Val int
//		     Left *TreeNode
//		     Right *TreeNode
//		 }
func bstToGst(root *TreeNode) *TreeNode {
	var sum int
	bypass(root, &sum)
	return root
}
func bypass(leaf *TreeNode, sum *int) {
	if leaf.Right != nil {
		bypass(leaf.Right, sum)
	}
	*sum += leaf.Val
	leaf.Val = *sum
	if leaf.Left != nil {
		bypass(leaf.Left, sum)
	}
}

//----------------------------------------------------5
// type ListNode struct {
//	     Val int
//	     Next *ListNode
//	 }

func pairSum(head *ListNode) int {
	l := lenNodes(head)
	var resSum int
	countF, countB := 0, l-1
	for i := 0; i < (l / 2); i++ {
		nHead, newSum := head, 0
		for q := 0; q < i; q++ {
			nHead = nHead.Next
		}
		for j := i; j < l-i; j++ {
			if j == countF || j == countB {
				newSum += nHead.Val
			}
			nHead = nHead.Next
		}
		if resSum < newSum {
			resSum = newSum
		}
		countF += 1
		countB -= 1
	}
	return resSum
}
func lenNodes(head *ListNode) int {
	var count int
	cHead := head
	for cHead != nil {
		count++
		cHead = cHead.Next
	}
	return count
}

func addNode(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	if head == nil {
		head = newNode
	} else {
		curNode := head
		for curNode.Next != nil {
			curNode = curNode.Next
		}
		curNode.Next = newNode
	}
	return head
}

//func pairSum(head *ListNode) int {
//	var array = make([]int, 0)
//	current := head
//	for current != nil {
//		array = append(array, current.Val)
//		current = current.Next
//	}
//	var max = array[0]
//
//	if len(array) > 0 {
//		for i := 0; i < len(array)/2+1; i++ {
//			twin := len(array) - 1 - i
//			if array[i]+array[twin] > max {
//				max = array[i] + array[twin]
//			}
//		}
//	}
//	return max
//}
// ---------------------------------------------------------6
//type TreeNode struct {
//	    Val int
//	    Left *TreeNode
//		Right *TreeNode
//	}

func constructMaximumBinaryTree(root []int) *TreeNode {
	if len(root) == 1 {
		return &TreeNode{Val: root[0]}
	}
	top := root[0]
	topI := 0
	for i := 1; i < len(root); i++ {
		if top < root[i] {
			top = root[i]
			topI = i
		}
	}
	leaf := TreeNode{Val: root[topI]}
	if topI+1 < len(root) {
		leaf.Right = constructMaximumBinaryTree(root[topI+1:])
	}
	if topI >= 1 {
		leaf.Left = constructMaximumBinaryTree(root[:topI])
	}
	return &leaf
}

// --------------------------------------------------7
//
//	type TreeNode struct {
//		    Val int
//		    Left *TreeNode
//		    Right *TreeNode
//		 }

func balanceBST(root *TreeNode) *TreeNode {
	var trees = []*TreeNode{}
	sortTree(root, &trees)
	return balance(trees)
}
func sortTree(node *TreeNode, nodes *[]*TreeNode) {
	if node == nil {
		return
	}
	if node.Left != nil {
		sortTree(node.Left, nodes)
	}
	*nodes = append(*nodes, node)
	if node.Right != nil {
		sortTree(node.Right, nodes)
	}
}
func balance(nod []*TreeNode) *TreeNode {
	lNod := len(nod)
	if lNod == 0 {
		return nil
	}
	if lNod == 1 {
		nod[0].Left, nod[0].Right = nil, nil
		return nod[0]
	}
	newTree := nod[lNod/2]
	newTree.Left = balance(nod[:lNod/2])
	newTree.Right = balance(nod[lNod/2+1:])
	return newTree
}

//-------------------------------------------------8

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	toiNodes := make([]int, n)
	for _, e := range edges {
		indexToi := e[1]
		toiNodes[indexToi]++
	}
	minReach := make([]int, 0, n)
	for i, d := range toiNodes {
		if d == 0 {
			minReach = append(minReach, i)
		}
	}
	return minReach
}

// -------------------------------------9
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var out = make([]bool, 0, len(l))
	for i, inxFirst := range l {
		var testList = make([]int, 0, len(nums))
		for j := inxFirst; j <= r[i]; j++ {
			testList = append(testList, nums[j])
		}
		sort.Slice(testList, func(i, j int) bool {
			return testList[i] < testList[j]
		})
		var boolTest = true
		for q := 1; q < len(testList); q++ {
			if testList[q]-testList[q-1] != testList[1]-testList[0] {
				boolTest = false
				break
			}
		}
		out = append(out, boolTest)
	}
	return out
}

// -----------------------------------------------10

func numTilePossibilities(tiles string) int {
	var tiLen = len(tiles)
	var count func(r string, size int)
	var seqs = map[string]bool{}
	var done = make([]bool, tiLen)
	var simb = strings.Split(tiles, "")

	count = func(r string, size int) {
		if len(r) == size {
			seqs[r] = true
			return
		}
		for i := 0; i < tiLen; i++ {
			if done[i] == false {
				done[i] = true
				count(r+simb[i], size)
				done[i] = false
			}
		}
		return
	}

	for size := 1; size <= tiLen; size++ {
		count("", size)
	}
	return len(seqs)
}

//--------------------------------------------------11
// type TreeNode struct {
//	     Val int
//	     Left *TreeNode
//	     Right *TreeNode
//	 }

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right = removeLeafNodes(root.Right, target)
	root.Left = removeLeafNodes(root.Left, target)
	if root.Val == target && root.Right == nil && root.Left == nil {
		return nil
	}
	return root
}

// --------------------------------------------------12
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	lFrom := list1
	for i := 0; i < a-1; i++ {
		lFrom = lFrom.Next
	}
	lTo := lFrom
	for i := a - 1; i <= b; i++ {
		lTo = lTo.Next
	}
	lFrom.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = lTo
	return list1
}

// ------------------------------------------13
func maxSum(grid [][]int) int {
	x, y := len(grid)-2, len(grid[0])-2
	var sum int
	for i := 0; i < x; i++ {
		var nextSum int
		for j := 0; j < y; j++ {
			nextSum = grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if sum < nextSum {
				sum = nextSum
			}
		}
	}
	return sum
}

// --------------------------------------------------14
func xorQueries(arr []int, queries [][]int) []int {
	if len(queries) == 0 || len(arr) == 0 {
		return []int{}
	}
	resL := make([]int, 0, len(queries))
	for _, val := range queries {
		if val[0] == val[1] {
			resL = append(resL, arr[val[0]])
			continue
		}
		var calc int
		for i := val[0]; i <= val[1]; i++ {
			if i == val[0] {
				calc = arr[i]
				continue
			}
			calc ^= arr[i]
		}
		resL = append(resL, calc)
	}
	return resL
}

// -----------------------------------------------15
func minPartitions(n string) int {
	numList, muxTry := strings.Split(n, ""), 0
	for _, val := range numList {
		valCur, _ := strconv.Atoi(val)
		if muxTry < valCur {
			muxTry = valCur
		}
	}
	return muxTry
}

// ------------------------------------------------------16

type SubrectangleQueries struct {
	rowCols [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rowCols: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rowCols[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rowCols[row][col]
}

// -------------------------------------17
func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))
	for i, que := range queries {
		for _, p := range points {
			if (que[0]-p[0])*(que[0]-p[0])+(que[1]-p[1])*(que[1]-p[1]) <= que[2]*que[2] {
				res[i] += 1
			}
		}
	}
	return res
}

// -----------------------------------------18
func groupThePeople(groupSizes []int) [][]int {
	var result = make([][]int, 0, len(groupSizes))
	var chackMap = make(map[int][]int, len(groupSizes))

	for i, s := range groupSizes {
		chackMap[s] = append(chackMap[s], i)
		if len(chackMap[s]) == s {
			result = append(result, chackMap[s])
			delete(chackMap, s)
		}
	}
	for _, g := range chackMap {
		result = append(result, g)
	}

	return result
}

//-----------------------------------------------------19
// type TreeNode struct {
//	     Val int
//	     Left *TreeNode
//	     Right *TreeNode
//	 }

func averageOfSubtree(root *TreeNode) int {
	_, _, result := findNodes(root)
	return result
}

func findNodes(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	sum1, count1, middle1 := findNodes(root.Left)
	sum2, count2, middle2 := findNodes(root.Right)
	sum := sum1 + sum2 + root.Val
	count := count1 + count2 + 1
	midSum := middle1 + middle2
	if root.Val == sum/count {
		midSum++
	}
	return sum, count, midSum
}

// -----------------------------------------------------20
func processQueries(queries []int, m int) []int {
	var res = make([]int, len(queries))
	var newQ = make([]int, m)
	for i := 0; i < m; i++ {
		newQ[i] = i + 1
	}
	for i, v := range queries {
		j := 0
		for newQ[j] != v {
			j++
		}
		res[i] = j
		for ; j > 0; j-- {
			newQ[j] = newQ[j-1]
		}
		newQ[0] = v
	}
	return res
}
func main() {
	//1
	//fmt.Println(deepestLeavesSum(&TreeNode{1, 2, 3, 4, 5, null, 6, 7, null, null, null, null, 8}))
	//2
	//fmt.Println(sortTheStudents([][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}, 2))
	//3
	// 4
	//5
	//node1 := &ListNode{Val: 7}
	//testL := []int{57, 13, 31, 17, 65, 32, 3, 97, 22, 7, 20, 69, 35, 69, 75, 13, 33, 50, 80, 64, 71, 15, 28, 2, 27, 39, 48, 13, 22, 84, 5, 51, 46, 26, 78, 56, 63}
	//for _, val := range testL {
	//	addNode(node1, val)
	//}
	//sumL := pairSum(node1)
	//fmt.Println(sumL)
	//6
	//7
	//8
	//fmt.Println(findSmallestSetOfVertices(5, [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}))
	//9
	//fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
	//10
	//fmt.Println(numTilePossibilities("AAB"))
	//11

	//-----------------------------------12
	//node1 := &ListNode{Val: 90}
	//testL := []int{91, 92, 93, 94, 95}
	//for _, val := range testL {
	//	addNode(node1, val)
	//}
	//node2 := &ListNode{Val: 4}
	//testL2 := []int{5, 6}
	//for _, val := range testL2 {
	//	addNode(node2, val)
	//}
	//mergeInBetween(node1, 2, 4, node2)
	//n := node1
	//for n != nil {
	//	fmt.Println(n.Val)
	//	n = n.Next
	//}
	//--------------------------------------------13
	//fmt.Println(maxSum([][]int{{520626, 685427, 788912, 800638, 717251, 683428}, {23602, 608915, 697585, 957500, 154778, 209236}, {287585, 588801, 818234, 73530, 939116, 252369}}))
	//14
	//fmt.Println(xorQueries([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}))
	//15
	//fmt.Println(minPartitions("82734"))
	//16
	//17
	//fmt.Println(countPoints([][]int{{1, 3}, {3, 3}, {5, 3}, {2, 2}}, [][]int{{2, 3, 1}, {4, 3, 1}, {1, 1, 2}}))
	//18
	//fmt.Println(groupThePeople([]int{3, 3, 3, 3, 3, 1, 3}))
	//19
	//20
	fmt.Println(processQueries([]int{3, 1, 2, 1}, 5))
}
