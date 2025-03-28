package main

import (
	"testing"
)

func TestBinarySearchSqrt(t *testing.T) {
	tests := []struct {
		target   int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{8, 3},
		{9, 3},
		{10, 3},
		{11, 3},
		{15, 4},
		{16, 4},
		{17, 4},
		{23, 5},
		{24, 5},
		{25, 5},
		{26, 5},
	}

	for _, test := range tests {
		ans := binarySearchSqrt(test.target)

		if ans != test.expected {
			t.Errorf("input %v, expected %v, got %v", test.target, test.expected, ans)
		}

	}
}
