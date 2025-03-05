package main

import (
	"testing"
	"vkalgo/dz2/linkedlist"
)

func TestRemove(t *testing.T) {
	tests := []struct {
		list     *linkedlist.List
		target   int
		expected *linkedlist.Node
	}{
		{
			list:     &linkedlist.List{},
			target:   1,
			expected: nil,
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(3)
				l.AppendFront(2)
				l.AppendFront(1)
				return l
			}(),
			target:   2,
			expected: &linkedlist.Node{Data: 1, Next: &linkedlist.Node{Data: 3, Next: nil}},
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(3)
				l.AppendFront(2)
				l.AppendFront(1)
				return l
			}(),
			target:   1,
			expected: &linkedlist.Node{Data: 2, Next: &linkedlist.Node{Data: 3, Next: nil}},
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(3)
				l.AppendFront(2)
				l.AppendFront(1)
				return l
			}(),
			target:   4,
			expected: &linkedlist.Node{Data: 1, Next: &linkedlist.Node{Data: 2, Next: &linkedlist.Node{Data: 3, Next: nil}}},
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(1)
				return l
			}(),
			target:   1,
			expected: nil,
		},
	}

	for _, tt := range tests {
		ans := Solve(tt.list, tt.target)

		if !compareNodes(ans, tt.expected) {
			tt.list.PrintList()
			t.Errorf(" expected %v, but got %v", tt.expected, ans)
		}
	}

}

func compareNodes(n1, n2 *linkedlist.Node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}

	if n1.Data != n2.Data {
		return false
	}
	return compareNodes(n1.Next, n2.Next)
}
