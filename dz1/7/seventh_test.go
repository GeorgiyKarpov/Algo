package main

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		a        []int
		expected []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{1, 0, 1, 2, 2, 0, 2, 1}, []int{0, 0, 1, 1, 1, 2, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
	}

	for _, tt := range tests {
		ans := Solve(tt.a)
		if !reflect.DeepEqual(ans, tt.expected) {
			t.Errorf("input %v, expected %v, but got %v", tt.a, tt.expected, ans)
		}
	}
}
