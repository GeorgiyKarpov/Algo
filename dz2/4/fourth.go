package main

import "vkalgo/dz2/linkedlist"

func Solve(a *linkedlist.List, target int) *linkedlist.Node {
	if a.Head == nil {
		return nil
	}

	dummy := &linkedlist.Node{Data: 0, Next: a.Head}

	cur := dummy
	next := cur.Next
	for next != nil {
		if next.Data == target {
			cur.Next = next.Next
			break
		}
		cur = cur.Next
		next = next.Next
	}
	a.Head = dummy.Next
	return a.Head
}
