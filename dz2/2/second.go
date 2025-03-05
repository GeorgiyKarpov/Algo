package main

import "vkalgo/dz2/linkedlist"

func Solve(a *linkedlist.List) *linkedlist.Node {
	cur := a.Head
	var prev *linkedlist.Node
	var next *linkedlist.Node
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	a.Head = prev

	return a.Head
}

func main() {
	a := &linkedlist.List{}
	Solve(a)
}
