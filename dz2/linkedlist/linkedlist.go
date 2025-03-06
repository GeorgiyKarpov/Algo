package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) AppendFront(n int) {
	if l.Head == nil {
		newNode := &Node{Data: n, Next: nil}
		l.Head = newNode
		return
	}
	newNode := &Node{Data: n, Next: l.Head}
	l.Head = newNode
}

func (l *List) AppendBack(n int) {
	newNode := &Node{Data: n, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return
	}

	cur_node := l.Head
	for cur_node.Next != nil {
		cur_node = cur_node.Next
	}

	cur_node.Next = newNode
}

func (l *List) InsertAfterValue(after, n int) {
	search := l.Head
	for search != nil {
		if search.Data == after {
			break
		}
		search = search.Next
	}

	if search != nil {
		newNode := &Node{Data: n, Next: search.Next}
		search.Next = newNode
		return
	}
	fmt.Printf("Cannot insert node, after value %d doesn't exist", after)
	fmt.Println()
}

func (l *List) PrintList() {
	cur_node := l.Head
	output := ""
	for cur_node != nil {
		output += fmt.Sprintf("%d ", cur_node.Data)
		if cur_node.Next != nil {
			output += " -> "
		}
		cur_node = cur_node.Next
	}
	fmt.Println(output)
}

func (l *List) CreateCycle(at int) {
	if l.Head == nil {
		return
	}

	current := l.Head
	for i := 0; i < at; i++ {
		if current == nil {
			return
		}
		current = current.Next
	}
	if current == nil {
		return
	}

	tail := l.Head
	for tail.Next != nil {
		tail = tail.Next
	}
	if tail != current {
		tail.Next = current
	}
}
