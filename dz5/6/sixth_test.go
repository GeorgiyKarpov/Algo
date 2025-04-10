package main

import "testing"

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		subRoot  *TreeNode
		expected bool
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 5},
			},
			subRoot: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: true,
		},
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 0},
					},
				},
				Right: &TreeNode{Val: 5},
			},
			subRoot: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			root:     nil,
			subRoot:  nil,
			expected: true,
		},
		{
			root:     &TreeNode{Val: 1},
			subRoot:  nil,
			expected: true,
		},
		{
			root:     nil,
			subRoot:  &TreeNode{Val: 1},
			expected: false,
		},
		{
			root:     &TreeNode{Val: 1},
			subRoot:  &TreeNode{Val: 1},
			expected: true,
		},
	}

	for _, test := range tests {
		ans := isSubtree(test.root, test.subRoot)
		if ans != test.expected {
			t.Errorf("Input %v %v got %v, but expected %v", test.root, test.subRoot, ans, test.expected)
		}
	}
}
