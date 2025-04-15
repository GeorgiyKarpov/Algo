package main

func isMaxHeap(arr []int) bool {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		left := 2*i + 1
		right := 2*i + 2
		if left < n && arr[i] < arr[left] {
			return false
		}
		if right < n && arr[i] < arr[right] {
			return false
		}
	}
	return true
}
