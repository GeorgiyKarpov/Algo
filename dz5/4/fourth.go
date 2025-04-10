package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxMinMultiplication(data []int) int {
	if len(data) == 1 {
		return data[0]
	}
	if len(data) == 2 {
		return data[0] * data[1]
	}

	minIdx := 1
	maxIdx := 2

	for i := 1; i < len(data); i = 2*i + 1 {
		minIdx = i
	}

	for i := 2; i < len(data); i = 2*i + 2 {
		maxIdx = i
	}

	return data[minIdx] * data[maxIdx]

}
