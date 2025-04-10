package main

import "testing"

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		root     *TreeNode
		expected int
	}{
		{
			root:     nil,
			expected: 0,
		},
		{
			root:     &TreeNode{Val: 1},
			expected: 0,
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
			expected: 3,
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
			expected: 1,
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
			expected: 0,
		},
	}

	for _, test := range tests {
		ans := countMirrorTwinsDfs(test.root)
		if ans != test.expected {
			t.Errorf("input %v expected %v, but got %v", test.root, test.expected, ans)
		}
	}
}
