package main

import (
	"reflect"
	"testing"
)

func TestPartReverse(t *testing.T) {
	tests := []struct {
		a        []int
		k        int
		expected []int
	}{
		{[]int{}, 3, []int{}},
		{[]int{1}, 5, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 8, []int{7, 1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
	}

	for _, tt := range tests {
		ans := Solve(tt.a, len(tt.a), tt.k)

		if !reflect.DeepEqual(ans, tt.expected) {
			t.Errorf("input %v, %d, expected %v, got %v", tt.a, tt.k, tt.expected, ans)
		}
	}
}
