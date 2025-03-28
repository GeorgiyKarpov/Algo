package main

import "testing"

func TestFeedAnimals(t *testing.T) {
	tests := []struct {
		animals  []int
		food     []int
		expected int
	}{
		{
			animals:  []int{},
			food:     []int{},
			expected: 0,
		},
		{
			animals:  []int{},
			food:     []int{1, 2, 3},
			expected: 0,
		},
		{
			animals:  []int{1, 2, 3},
			food:     []int{},
			expected: 0,
		},
		{
			animals:  []int{4},
			food:     []int{7},
			expected: 1,
		},
		{
			animals:  []int{2},
			food:     []int{1},
			expected: 0,
		},
		{
			animals:  []int{3, 4, 7},
			food:     []int{8, 1, 2},
			expected: 1,
		},
		{
			animals:  []int{3, 8, 1, 4},
			food:     []int{1, 1, 2},
			expected: 1,
		},
		{
			animals:  []int{1, 2, 2},
			food:     []int{7, 1},
			expected: 2,
		},
		{
			animals:  []int{8, 2, 3, 2},
			food:     []int{1, 4, 3, 8},
			expected: 3,
		},
		{
			animals:  []int{7, 6, 7},
			food:     []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			animals:  []int{2, 2, 2},
			food:     []int{2, 2, 1, 5},
			expected: 3,
		},
		{
			animals:  []int{1, 2, 4, 100},
			food:     []int{100, 1, 3, 3},
			expected: 3,
		},
	}

	for _, test := range tests {
		ans := feedAnimals(test.animals, test.food)
		if ans != test.expected {
			t.Errorf("input %v %v, expected %v, but got %v", test.animals, test.food, test.expected, ans)
		}
	}
}
