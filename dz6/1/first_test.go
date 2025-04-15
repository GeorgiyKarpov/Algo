package main

import "testing"

func buildTree(arr []int, i int) *TreeNode {
	if i >= len(arr) {
		return nil
	}
	node := &TreeNode{Val: arr[i]}
	node.Left = buildTree(arr, 2*i+1)
	node.Right = buildTree(arr, 2*i+2)
	return node
}

func TestIsMaxHeap(t *testing.T) {
	tests := []struct {
		arr      []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{42}, true},
		{[]int{10, 5}, true},
		{[]int{5, 10}, false},
		{[]int{10, 9, 8, 7, 6, 5, 4}, true},
		{[]int{10, 15, 8, 7, 6, 13, 1}, false},
		{[]int{100, 45, 32, 14, 7, 1}, true},
	}

	for _, test := range tests {
		ans := isMaxHeap(test.arr)
		if ans != test.expected {
			t.Errorf("Input %v, expected %v, but got %v", test.arr, test.expected, ans)
		}
	}

	for _, test := range tests {
		root := buildTree(test.arr, 0)
		ans := isMaxHeapBfs(root)
		if ans != test.expected {
			t.Errorf("isMaxHeapBfs: Input %v, expected %v, but got %v", test.arr, test.expected, ans)
		}
	}

}
