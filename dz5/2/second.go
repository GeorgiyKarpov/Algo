package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			if queue[i] == nil && queue[n-i-1] == nil {
				continue
			}
			if queue[i] == nil || queue[n-i-1] == nil {
				return false
			}
			if queue[i].Val != queue[n-i-1].Val {
				return false
			}
			queue = append(queue, queue[i].Left, queue[i].Right)
		}
		queue = queue[n:]
	}
	return true
}
