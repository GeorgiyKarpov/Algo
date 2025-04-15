package main

import "container/heap"

type Elem struct {
	Val    int
	ArrIdx int
	Idx    int
}

type MinHeap []Elem

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Elem))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	elem := old[n-1]
	*h = old[0 : n-1]
	return elem
}

func mergeKSortedArrays(sortedArrays [][]int) []int {
	var mergedArray []int
	h := &MinHeap{}
	heap.Init(h)

	for i, a := range sortedArrays {
		if len(a) > 0 {
			heap.Push(h, Elem{
				Val:    a[0],
				ArrIdx: i,
				Idx:    0,
			})
		}
	}

	for h.Len() > 0 {
		elem := heap.Pop(h).(Elem)
		mergedArray = append(mergedArray, elem.Val)

		if elem.Idx+1 < len(sortedArrays[elem.ArrIdx]) {
			heap.Push(h, Elem{
				Val:    sortedArrays[elem.ArrIdx][elem.Idx+1],
				ArrIdx: elem.ArrIdx,
				Idx:    elem.Idx + 1,
			})
		}
	}

	return mergedArray
}
