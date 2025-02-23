package main

import (
	"reflect"
	"testing"
)

func TestEvenFirst(t *testing.T) {
	tests := []struct {
		a        []int
		expected []int
	}{
		{[]int{3, 2, 4, 1, 11, 8, 9}, []int{2, 4, 8, 1, 11, 3, 9}},
		{[]int{2, 2, 4, 6, 7, 8}, []int{2, 2, 4, 6, 8, 7}},
		{[]int{83, 17, 91, 2, 6, 4}, []int{2, 6, 4, 93, 17, 91}},
	}

	for _, tt := range tests {
		ans := Solve(tt.a)

		if !reflect.DeepEqual(ans, tt.a) {
			t.Errorf("input %v, expected %v, but got %v", tt.a, tt.expected, ans)
		}
	}
}
