package main

import (
	"reflect"
	"testing"
	"vkalgo/dz2/linkedlist"
)

func createList(data []int) *linkedlist.Node {
	l := &linkedlist.List{}
	for _, v := range data {
		l.AppendBack(v)
	}
	return l.Head
}

func listToSlice(head *linkedlist.Node) []int {
	var result []int
	for head != nil {
		result = append(result, head.Data)
		head = head.Next
	}
	return result
}

func TestSolve(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{

		{nil, nil, nil},
		{nil, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2}, nil, []int{1, 2}},
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}},
		{[]int{1, 3, 5}, []int{0, 2, 4, 6, 8}, []int{0, 1, 2, 3, 4, 5, 6, 8}},
	}

	for _, tt := range tests {
		l1 := createList(tt.list1)
		l2 := createList(tt.list2)
		merged := Solve(l1, l2)
		result := listToSlice(merged)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("input %v %v got %v but expected %v", tt.list1, tt.list2, result, tt.expected)
		}
	}
}
