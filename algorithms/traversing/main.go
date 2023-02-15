package main

import (
	"fmt"
	"traversing/leetcode"
)

// DepthFirstSearch - поиск в глубину (ищет по "вертикальным уровням")
func DepthFirstSearch(root *leetcode.TreeNode) []int {
	if root == nil { // Если вершины не существует, возвращаем пустой результат
		return []int{}
	}

	searchStack := []*leetcode.TreeNode{root} // Для обхода используется СТЕК
	visited := make(map[*leetcode.TreeNode]bool)

	var res []int

	for len(searchStack) > 0 {
		node := searchStack[len(searchStack)-1]        // Достаём с вершины стека
		searchStack = searchStack[:len(searchStack)-1] // Соответствено, нужно убрать полученный элемент с вершины стека

		if node != nil && !visited[node] { // Если элемент ещё не посещался и существует
			visited[node] = true        // Помечаем элемент как посещённый
			res = append(res, node.Val) // Добавляем посещённый только что элемент к результату

			searchStack = append(searchStack, node.Right, node.Left)
		}
	}

	return res
}

// BreathFirstSearch - поиск в ширину (ищет по "горизонтальным уровням")
func BreathFirstSearch(root *leetcode.TreeNode) []int {
	if root == nil { // Если вершины не существует, возвращаем пустой результат
		return []int{}
	}

	searchQueue := []*leetcode.TreeNode{root} // Для обхода используется ОЧЕРЕДЬ
	visited := make(map[*leetcode.TreeNode]bool)

	var res []int

	for len(searchQueue) > 0 {
		node := searchQueue[0]        // Достаём из очереди
		searchQueue = searchQueue[1:] // После получеения элемента нужно уменьшить очередь

		if node != nil && !visited[node] {
			visited[node] = true        // Помечаем элемент как посещённый
			res = append(res, node.Val) // Добавляем посещённый только что элемент к результату

			searchQueue = append(searchQueue, node.Left, node.Right)
		}
	}

	return res
}

const null = leetcode.NULL

func main() {
	root := leetcode.NewTreeNode(3, 9, 10, null, null, 5, 7, 2, null, 1, 4, null, null, null, 6)

	root.Print()

	fmt.Printf("DFS: %v\n", DepthFirstSearch(root))
	fmt.Printf("BFS: %v\n", BreathFirstSearch(root))
}
