package main

type TreeNode struct {
	Val     int
	Left    *TreeNode
	Right   *TreeNode
	Balance int
}

func calculateHeightsAndBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftH := calculateHeightsAndBalance(node.Left)
	rightH := calculateHeightsAndBalance(node.Right)

	node.Balance = leftH - rightH

	return 1 + max(leftH, rightH)
}
