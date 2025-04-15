package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	leaf := false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			if leaf == true {
				return false
			}
			queue = append(queue, node.Left)
		} else {
			leaf = true
		}

		if node.Right != nil {
			if leaf == true {
				return false
			}
			queue = append(queue, node.Right)
		} else {
			leaf = true
		}
	}

	return true
}
