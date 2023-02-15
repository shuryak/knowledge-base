package leetcode

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NULL = math.MinInt

func (node *TreeNode) Print() {
	maxLevel := node.maxLevel()

	node.printNode([]*TreeNode{node}, 1, maxLevel)
}

func (node *TreeNode) printNode(nodes []*TreeNode, level, maxLevel int) {
	if len(nodes) == 0 || node.isAllElementsNil(nodes) {
		return
	}

	floor := maxLevel - level
	endgeLines := int(math.Pow(2, math.Max(float64(floor-1), 0)))
	firstSpaces := int(math.Pow(2, float64(floor)) - 1)
	betweenSpaces := int(math.Pow(2, float64(floor+1)) - 1)

	node.printWhitespaces(firstSpaces)

	var newNodes []*TreeNode
	for _, node := range nodes {
		if node != nil {
			fmt.Print(node.Val)
			newNodes = append(newNodes, node.Left, node.Right)
		} else {
			newNodes = append(newNodes, nil, nil)
			fmt.Print(" ")
		}

		node.printWhitespaces(betweenSpaces)
	}
	fmt.Println()

	for i := 1; i < endgeLines; i++ {
		for j := 0; j < len(nodes); j++ {
			node.printWhitespaces(firstSpaces - i)
			if nodes[j] == nil {
				node.printWhitespaces(endgeLines + endgeLines + i + 1)
				continue
			}

			if nodes[j].Left != nil {
				fmt.Print("/")
			} else {
				node.printWhitespaces(1)
			}

			node.printWhitespaces(i + i - 1)

			if nodes[j].Right != nil {
				fmt.Printf("\\")
			} else {
				node.printWhitespaces(1)
			}

			node.printWhitespaces(endgeLines + endgeLines - i)
		}

		fmt.Println()
	}

	node.printNode(newNodes, level+1, maxLevel)
}

func (node *TreeNode) printWhitespaces(count int) {
	for i := 0; i < count; i++ {
		fmt.Print(" ")
	}
}

func (node *TreeNode) maxLevel() int {
	if node == nil {
		return 0
	}

	return 1 + int(math.Max(float64(node.Left.maxLevel()), float64(node.Right.maxLevel())))
}

func (node *TreeNode) isAllElementsNil(nodes []*TreeNode) bool {
	for _, node := range nodes {
		if node != nil {
			return false
		}
	}

	return true
}

func NewTreeNode(values ...int) *TreeNode {
	n := len(values)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: values[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && values[i] != NULL {
			node.Left = &TreeNode{Val: values[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && values[i] != NULL {
			node.Right = &TreeNode{Val: values[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
