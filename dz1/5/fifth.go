package main

import "fmt"

func Solve(a1, a2 []int, m, n int) []int {
	if n == 0 {
		return a1
	}
	if m == 0 {
		return a2
	}
	p1 := m - 1
	p2 := n - 1
	p3 := m + n - 1

	for p3 >= 0 {
		if p1 >= 0 && a1[p1] >= a2[p2] {
			a1[p3] = a1[p1]
			p1--
		} else {
			a1[p3] = a2[p2]
			p2--
			if p2 < 0 {
				break
			}
		}
		p3--
	}

	return a1
}

func main() {
	var m, n int
	fmt.Scan(&m)
	fmt.Scan(&n)
	a1 := make([]int, m+n)
	for i := 0; i < m; i++ {
		fmt.Scan(&a1[i])
	}
	a2 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a2[i])
	}

	ans := Solve(a1, a2, m, n)
	fmt.Println(ans)
}
