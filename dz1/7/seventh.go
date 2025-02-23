package main

import "fmt"

func Solve(a []int) []int {
	p0 := 0
	mid := 0
	p2 := len(a) - 1
	for mid <= p2 {
		if a[mid] == 2 {
			a[mid], a[p2] = a[p2], a[mid]
			p2--
		} else if a[mid] == 0 {
			a[mid], a[p0] = a[p0], a[mid]
			p0++
			mid++
		} else {
			mid++
		}
	}
	return a
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(Solve(a))
}
