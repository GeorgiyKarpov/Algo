package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		data     []int
		target   int
		expected []int
	}{
		{
			data:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			data:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			data:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			data:     []int{1, 2, 3},
			target:   6,
			expected: []int{},
		},
		{
			data:     []int{1},
			target:   2,
			expected: []int{},
		},
	}

	for _, test := range tests {
		ans := twoSum(test.data, test.target)
		if len(ans) != len(test.expected) {
			t.Errorf("input %v %v, expected %v, but got %v", test.data, test.target, test.expected, ans)
		} else if !reflect.DeepEqual(ans, test.expected) {
			t.Errorf("input %v %v, expected %v, but got %v", test.data, test.target, test.expected, ans)
		}
	}
}
