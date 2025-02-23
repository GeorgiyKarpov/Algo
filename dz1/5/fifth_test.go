package main

import (
	"reflect"
	"testing"
)

func TestMergeArrayModify(t *testing.T) {
	tests := []struct {
		m        int
		a1       []int
		n        int
		a2       []int
		expected []int
	}{
		{2, []int{1, 2}, 0, []int{}, []int{1, 2}},
		{0, []int{0, 0, 0}, 3, []int{1, 2, 3}, []int{1, 2, 3}},
		{4, []int{3, 8, 10, 11, 0, 0, 0}, 3, []int{1, 7, 9}, []int{1, 3, 7, 8, 9, 10, 11}},
		{3, []int{0, 0, 0, 0, 0, 0}, 3, []int{0, 2, 3}, []int{0, 0, 0, 0, 2, 3}},
	}

	for _, tt := range tests {
		ans := Solve(tt.a1, tt.a2, tt.m, tt.n)

		if !reflect.DeepEqual(ans, tt.expected) {
			t.Errorf("input %v, %v, expected %v, but got %v", tt.a1, tt.a2, tt.expected, ans)
		}
	}
}
