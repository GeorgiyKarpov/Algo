package main

func twoSum(data []int, target int) []int {
	a := make(map[int]int)
	for i, j := range data {
		a[j] = i
	}
	ans := make([]int, 0)
	for i := 0; i < len(data); i++ {
		if j, ok := a[target-data[i]]; ok && j != i {
			ans = append(ans, i, a[target-data[i]])
			break
		}
	}
	return ans
}
