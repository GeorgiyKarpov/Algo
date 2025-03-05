package main

import (
	"testing"
	"vkalgo/dz2/linkedlist"
)

func TestMid(t *testing.T) {
	tests := []struct {
		list     *linkedlist.List
		expected *linkedlist.Node
	}{
		{
			list:     &linkedlist.List{},
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
			expected: &linkedlist.Node{Data: 2, Next: &linkedlist.Node{Data: 3, Next: nil}},
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(3)
				return l
			}(),
			expected: &linkedlist.Node{Data: 3, Next: nil},
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(17)
				l.AppendFront(16)
				l.AppendFront(11)
				l.AppendFront(6)
				l.AppendFront(9)
				l.AppendFront(8)
				l.AppendFront(3)
				return l
			}(),
			expected: &linkedlist.Node{Data: 6, Next: &linkedlist.Node{Data: 11, Next: &linkedlist.Node{Data: 16, Next: &linkedlist.Node{Data: 17, Next: nil}}}}},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(4)
				l.AppendFront(3)
				l.AppendFront(2)
				l.AppendFront(1)
				return l
			}(),
			expected: &linkedlist.Node{Data: 2, Next: &linkedlist.Node{Data: 3, Next: &linkedlist.Node{Data: 4, Next: nil}}},
		},
	}

	for _, tt := range tests {
		ans := Solve(tt.list)
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
