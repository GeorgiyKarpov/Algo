package main

import (
	"reflect"
	"testing"
)

func copyTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  copyTree(root.Left),
		Right: copyTree(root.Right),
	}
}

func TestMirrorTree(t *testing.T) {
	tests := []struct {
		p        *TreeNode
		expected *TreeNode
	}{
		{nil, nil},
		{&TreeNode{Val: 1}, &TreeNode{Val: 1}},
		{
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			&TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 2},
			},
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 7},
				},
			},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 4},
				},
			},
		},
	}

	for _, test := range tests {
		ans := mirrorTree(copyTree(test.p))
		if !reflect.DeepEqual(ans, test.expected) {
			t.Errorf("input %v expected %v, but got %v", test.p, test.expected, ans)
		}
	}

	for _, test := range tests {
		ans := mirrorTreeBFS(test.p)
		if !reflect.DeepEqual(ans, test.expected) {
			t.Errorf("BFS input %v expected %v, but got %v", test.p, test.expected, ans)
		}
	}
}
