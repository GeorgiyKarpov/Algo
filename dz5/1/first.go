package main

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(arr []int, i int) *TreeNode {
	if i >= len(arr) {
		return nil
	}
	node := &TreeNode{Data: arr[i]}
	node.Left = buildTree(arr, 2*i+1)
	node.Right = buildTree(arr, 2*i+2)
	return node
}
