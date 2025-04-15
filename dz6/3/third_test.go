package main

import (
	"reflect"
	"testing"
)

func TestMergeKSortedArrays(t *testing.T) {
	tests := []struct {
		a        [][]int
		expected []int
	}{
		{
			a:        [][]int{},
			expected: []int{},
		},
		{
			a:        [][]int{{}, {}},
			expected: []int{},
		},
		{
			a:        [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			a: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
	}

	for _, test := range tests {
		ans := mergeKSortedArrays(test.a)
		if !reflect.DeepEqual(ans, test.expected) {
			t.Errorf("Input %v, expected %v, but got %v", test.a, test.expected, ans)
		}
	}
}
