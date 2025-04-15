package main

import "testing"

func TestIsCompleteTree(t *testing.T) {
	tests := []struct {
		p        *TreeNode
		k        int
		expected int
	}{
		{
			p: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{Val: 4},
			},
			k:        1,
			expected: 1,
		},
		{
			p: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 1,
						},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:        3,
			expected: 3,
		},
	}

	for _, test := range tests {
		ans := kthSmallest(test.p, test.k)
		if ans != test.expected {
			t.Errorf("Input %v, expected %v, but got %v", test.p, ans, test.expected)
		}
	}
}
