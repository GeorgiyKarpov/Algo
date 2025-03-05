package main

import "vkalgo/dz2/linkedlist"

func Solve(list1, list2 *linkedlist.Node) *linkedlist.Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &linkedlist.Node{}
	cur := dummy

	for list1 != nil && list2 != nil {
		if list1.Data >= list2.Data {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}

	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}
	return dummy.Next
}
