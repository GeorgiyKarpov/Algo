package main

import (
	"reflect"
	"testing"
)

func bfs(root *TreeNode) []int {
	var result []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		result = append(result, node.Data)
		queue = append(queue, node.Left, node.Right)
	}
	return result
}

func TestBuildTree(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{},
			expected: []int{},
		},
		{
			arr:      []int{1},
			expected: []int{1},
		},
		{
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			arr:      []int{8, 8, 14, 2, 7},
			expected: []int{8, 8, 14, 2, 7},
		},
	}

	for _, test := range tests {
		ans := bfs(buildTree(test.arr, 0))
		if !reflect.DeepEqual(ans, test.expected) {
			t.Errorf("input %v got = %v, but expected %v", test.arr, ans, test.expected)
		}
	}
}
