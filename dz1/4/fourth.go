package main

import "fmt"

func Solve(a1, a2, ans []int) []int {
	l := 0
	r := 0
	for l < len(a1) && r < len(a2) {
		if a1[l] <= a2[r] {
			ans = append(ans, a1[l])
			l++
		} else {
			ans = append(ans, a2[r])
			r++
		}
	}
	if l != len(a1) {
		ans = append(ans, a1[l:]...)
	} else if r != len(a2) {
		ans = append(ans, a2[r:]...)
	}

	return ans
}

func main() {
	var n1, n2 int
	fmt.Scan(&n1, &n2)
	a1 := make([]int, n1)
	for i := 0; i < n1; i++ {
		fmt.Scan(&a1[i])
	}

	a2 := make([]int, n2)
	for i := 0; i < n2; i++ {
		fmt.Scan(&a2[i])
	}

	ans := make([]int, 0, n1+n2)
	ans = Solve(a1, a2, ans)
	fmt.Println(ans)
}
