package main

import "vkalgo/dz2/linkedlist"

func Solve(a *linkedlist.List) *linkedlist.Node {
	if a.Head == nil {
		return nil
	}
	slow := a.Head
	fast := slow.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func main() {
	a := &linkedlist.List{}
	Solve(a)
}
