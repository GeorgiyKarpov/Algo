package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(left, right *TreeNode) int {
	if left == nil || right == nil {
		return 0
	}
	count := 0
	if left.Val == right.Val {
		count = 1
	}
	count += dfs(left.Left, right.Right)
	count += dfs(left.Right, right.Left)
	return count
}

func countMirrorTwinsDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root.Left, root.Right)
}
