package main

import "testing"

func TestMinDepth(t *testing.T) {
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
			expected: 1,
		},
		{
			root: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			expected: 2,
		},
		{
			root: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
							Right: &TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			expected: 5,
		},
	}

	for _, test := range tests {
		ans := minDepth(test.root)
		if ans != test.expected {
			t.Errorf("input %v expected %v, but got %v", test.root, test.expected, ans)
		}
	}
}
