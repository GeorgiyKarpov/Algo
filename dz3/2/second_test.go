package main

import (
	"testing"
)

func TestCopyTime(t *testing.T) {
	tests := []struct {
		n, x, y  int
		expected int
	}{

		{n: 1, x: 1, y: 1, expected: 1},
		{n: 2, x: 1, y: 1, expected: 2},
		{n: 4, x: 1, y: 1, expected: 3},
		{n: 5, x: 1, y: 3, expected: 4},
		{n: 6, x: 2, y: 1, expected: 5},
		{n: 7, x: 3, y: 2, expected: 10},
	}

	for _, test := range tests {
		ans := copyTime(test.n, test.x, test.y)
		if ans != test.expected {
			t.Errorf("input %v %v %v, got %v, but expected %v", test.n, test.x, test.y, ans, test.expected)
		}
	}
}
