package main

import "testing"

func TestCalculateHeightsAndBalance(t *testing.T) {
	tests := []struct {
		p                *TreeNode
		expectedBalances map[int]int
	}{
		{
			p:                nil,
			expectedBalances: map[int]int{},
		},
		{
			p:                &TreeNode{Val: 1},
			expectedBalances: map[int]int{1: 0},
		},
		{
			p: &TreeNode{
				Val: 8,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 2,
					},
				},
			},
			expectedBalances: map[int]int{
				2: 0,
				7: 1,
				8: 2,
			},
		},
		{
			p: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 10},
			},
			expectedBalances: map[int]int{
				4:  0,
				10: 0,
				5:  0,
			},
		},
		{
			p: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{Val: 15},
			},
			expectedBalances: map[int]int{
				3:  0,
				5:  1,
				15: 0,
				10: 1,
			},
		},
	}

	for _, test := range tests {
		_ = calculateHeightsAndBalance(test.p)
		balance := make(map[int]int)
		var traverse func(*TreeNode)
		traverse = func(node *TreeNode) {
			if node == nil {
				return
			}
			balance[node.Val] = node.Balance
			traverse(node.Left)
			traverse(node.Right)
		}
		traverse(test.p)
		if len(balance) != len(test.expectedBalances) {
			t.Errorf("len balance != len expected")
		}
		for val, expectedBalance := range test.expectedBalances {
			if got, ok := balance[val]; !ok {
				t.Errorf("Input %v node %d not found", test.p, val)
			} else if got != expectedBalance {
				t.Errorf("Input %v node %d expected %d, but got %d", test.p, val, expectedBalance, got)
			}
		}
	}
}
