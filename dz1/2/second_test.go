package main

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		a        []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{3, 8, 6, 9, 9, 8, 6}, []int{6, 8, 9, 9, 6, 8, 3}},
	}

	for _, test := range tests {
		ans := Solve(test.a)

		if !reflect.DeepEqual(ans, test.expected) {
			t.Errorf("input %v, expected %v, got %v", test.a, test.expected, ans)
		}
	}
}
