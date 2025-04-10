package main

import "testing"

func TestMaxMinMultiplication(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		{data: []int{15}, expected: 15},
		{data: []int{2, 3}, expected: 6},
		{data: []int{6, 4, 8}, expected: 32},
		{data: []int{10, 5, 15, 2}, expected: 30},
		{data: []int{10, 5, 20, 1, 9}, expected: 20},
	}

	for _, test := range tests {
		ans := maxMinMultiplication(test.data)
		if ans != test.expected {
			t.Errorf("Input %v expected %v, but got %d", test.data, test.expected, ans)
		}
	}
}
