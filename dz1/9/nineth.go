package main

import "fmt"

func Solve(a []int) []int {
	p := 0
	for _, num := range a {
		if num != 0 {
			a[p] = num
			p++
		}
	}

	for p < len(a) {
		a[p] = 0
		p++
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
