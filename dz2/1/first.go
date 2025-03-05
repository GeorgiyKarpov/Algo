package main

import (
	"fmt"
	"vkalgo/dz2/linkedlist"
)

func Solve(l *linkedlist.List) bool {
	if l.Head == nil || l.Head.Next == nil {
		return false
	}
	slow := l.Head
	fast := l.Head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

func main() {
	a := &linkedlist.List{}
	if Solve(a) {
		fmt.Println("Yes")
	}
	fmt.Println("No cycles")
}
