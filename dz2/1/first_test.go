package main

import (
	"testing"
	"vkalgo/dz2/linkedlist"
)

func TestCycle(t *testing.T) {
	tests := []struct {
		list     *linkedlist.List
		expected bool
	}{
		{
			list:     &linkedlist.List{},
			expected: false,
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(3)
				l.AppendFront(2)
				l.AppendFront(1)
				return l
			}(),
			expected: false,
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(3)
				return l
			}(),
			expected: false,
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
				l.AppendFront(11)
				l.CreateCycle(2)
				return l
			}(),
			expected: true,
		},
		{
			list: func() *linkedlist.List {
				l := &linkedlist.List{}
				l.AppendFront(2)
				l.AppendFront(1)
				l.CreateCycle(0)
				return l
			}(),
			expected: true,
		},
	}

	for _, tt := range tests {
		ans := Solve(tt.list)
		if ans != tt.expected {
			t.Errorf("expected %v, but got %v", tt.expected, ans)
		}
	}

}
