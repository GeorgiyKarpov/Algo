package main

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p, q     *TreeNode
		expected bool
	}{
		{
			p:        nil,
			q:        nil,
			expected: true,
		},
		{
			p:        nil,
			q:        &TreeNode{Val: 1},
			expected: false,
		},
		{
			p:        &TreeNode{Val: 1},
			q:        &TreeNode{Val: 1},
			expected: true,
		},
		{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
			expected: true,
		},
		{
			p: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			},
			q: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
		{
			p: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 1},
			},
			q: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 2},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		ans := isSameTree(test.p, test.q)
		if ans != test.expected {
			t.Errorf("Input %v %v, expected %v, but got %v", test.p, test.q, ans, test.expected)
		}
	}
}
