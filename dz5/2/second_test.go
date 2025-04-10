package main

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected bool
	}{
		{
			root:     nil,
			expected: true,
		},
		{
			root:     &TreeNode{Val: 1},
			expected: true,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 3},
				},
			},
			expected: true,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			expected: false,
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:  4,
					Left: &TreeNode{Val: 7},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		ans := isSymmetric(test.root)
		if ans != test.expected {
			t.Errorf("input %v %v %v, expected %v, but got %v", test.root, test.root.Left, test.root.Right, test.expected, ans)
		}
	}
	for _, test := range tests {
		ans := isSymmetricDFS(test.root)
		if ans != test.expected {
			t.Errorf("DFS input %v %v %v, expected %v, but got %v", test.root, test.root.Left, test.root.Right, test.expected, ans)
		}
	}

}
