package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack []*TreeNode

func (s *Stack) Push(node *TreeNode) {
	*s = append(*s, node)
}

func (s *Stack) Pop() (*TreeNode, bool) {
	if len(*s) == 0 {
		return nil, false
	}
	i := len(*s) - 1
	node := (*s)[i]
	*s = (*s)[:i]
	return node, true
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func kthSmallest(root *TreeNode, k int) int {
	var s Stack
	count := 0
	node := root

	for node != nil || !s.IsEmpty() {
		for node != nil {
			s.Push(node)
			node = node.Left
		}
		node, _ = s.Pop()
		count++

		if count == k {
			return node.Val
		}

		node = node.Right
	}

	return -1
}
