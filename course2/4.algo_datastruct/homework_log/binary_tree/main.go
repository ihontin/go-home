package main

import "fmt"

//Пример 1. Использование дерева для организации иерархической структуры
//В этом примере мы создадим дерево для организации иерархической структуры каталогов и файлов в файловой системе.

//type Node struct {
//	Name     string
//	Children []*Node
//}
//
//func main() {
//	root := &Node{Name: "root"}
//
//	dir1 := &Node{Name: "dir1"}
//	dir2 := &Node{Name: "dir2"}
//
//	file1 := &Node{Name: "file1.txt"}
//	file2 := &Node{Name: "file2.txt"}
//
//	dir1.Children = append(dir1.Children, file1)
//	dir2.Children = append(dir2.Children, file2)
//
//	root.Children = append(root.Children, dir1, dir2)
//
//	printTree(root, 0)
//}
//
//func printTree(node *Node, level int) {
//	indent := ""
//	for i := 0; i < level; i++ {
//		indent += "  "
//	}
//
//	fmt.Println(indent + node.Name)
//
//	for _, child := range node.Children {
//		printTree(child, level+1)
//	}
//}

// Пример 2. Обход дерева в глубину (DFS)
// В этом примере мы реализуем алгоритм обхода дерева в глубину (DFS) с использованием рекурсии.
//type Node struct {
//	Value    int
//	Children []*Node
//}
//
//func main() {
//	root := &Node{Value: 1}
//
//	child1 := &Node{Value: 2}
//	child2 := &Node{Value: 3}
//	child3 := &Node{Value: 4}
//
//	root.Children = append(root.Children, child1, child2)
//	child2.Children = append(child2.Children, child3)
//
//	dfs(root)
//}
//
//func dfs(node *Node) {
//	fmt.Println(node.Value)
//
//	for _, child := range node.Children {
//		dfs(child)
//	}
//}

//Пример 3. Задача коммивояжера (TSP) на графе
//В этом примере мы решим задачу коммивояжера (TSP) на графе с использованием алгоритма поиска в глубину (DFS).

var minCost int

func main() {
	graph := [][]int{
		{0, 2, 9, 10},
		{1, 0, 6, 4},
		{15, 7, 0, 8},
		{6, 3, 12, 0},
	}

	visited := make([]bool, len(graph))
	path := make([]int, len(graph)+1)

	dfs(graph, visited, path, 0, 0, 1)

	fmt.Println("Минимальная стоимость:", minCost)
}

func dfs(graph [][]int, visited []bool, path []int, current, cost, level int) {
	if level == len(graph) {
		if graph[current][0] != 0 {
			cost += graph[current][0]
			if cost < minCost || minCost == 0 {
				minCost = cost
			}
		}
		return
	}

	visited[current] = true
	path[level] = current

	for i := 0; i < len(graph); i++ {
		if !visited[i] && graph[current][i] != 0 {
			dfs(graph, visited, path, i, cost+graph[current][i], level+1)
		}
	}

	visited[current] = false
}

//Это лишь несколько примеров использования деревьев и графов в языке программирования Golang.
//Деревья и графы предоставляют мощные инструменты для организации и обработки данных в различных задачах.
