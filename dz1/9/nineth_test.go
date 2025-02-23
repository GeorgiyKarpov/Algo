package main

import (
	"reflect"
	"testing"
)

func TestZeroToEnd(t *testing.T) {
	tests := []struct {
		a        []int
		expected []int
	}{
		{[]int{0, 0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0, 0}},
		{[]int{0, 33, 57, 88, 60, 0, 0, 80, 99}, []int{33, 57, 88, 60, 80, 99, 0, 0, 0}},
		{[]int{0, 0, 0, 18, 16, 0, 0, 77, 99}, []int{18, 16, 77, 99, 0, 0, 0, 0, 0}},
	}

	for _, tt := range tests {
		ans := Solve(tt.a)
		if !reflect.DeepEqual(ans, tt.a) {
			t.Errorf("input %v, expected %v, but got %v", tt.a, tt.expected, ans)
		}
	}
}
