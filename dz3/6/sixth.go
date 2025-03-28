package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	a := make(map[string][]string)
	for _, j := range strs {
		runes := []rune(j)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] > runes[j]
		})
		sorted := string(runes)
		a[sorted] = append(a[sorted], j)
	}
	ans := make([][]string, 0)
	for _, j := range a {
		ans = append(ans, j)
	}
	return ans
}
