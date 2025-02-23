package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoArray(t *testing.T) {
	tests := []struct {
		a1       []int
		a2       []int
		expected []int
	}{
		{[]int{}, []int{1}, []int{1}},
		{[]int{1}, []int{}, []int{1}},
		{[]int{3, 6, 10, 11}, []int{1, 7, 9}, []int{1, 3, 6, 7, 9, 10, 11}},
		{[]int{1, 7, 9}, []int{0, 0, 0, 0}, []int{0, 0, 0, 0, 1, 7, 9}},
	}

	for _, tt := range tests {
		ans := make([]int, 0, len(tt.a1)+len(tt.a2))
		ans = Solve(tt.a1, tt.a2, ans)

		if !reflect.DeepEqual(ans, tt.expected) {
			t.Errorf("input %v, %v,  expected %v, but got %v", tt.a1, tt.a2, tt.expected, ans)
		}
	}
}
