package main

import "testing"

func TestIsCompleteTree(t *testing.T) {
	tests := []struct {
		p        *TreeNode
		expected bool
	}{
		{
			p:        nil,
			expected: true,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 6},
				},
			},
			expected: true,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 7},
				},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		ans := isCompleteTree(test.p)
		if ans != test.expected {
			t.Errorf("Input %v, expected %v, but got %v", test.p, ans, test.expected)
		}
	}
}
