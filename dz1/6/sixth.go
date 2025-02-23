package main

import "fmt"

func Solve(a []int) []int {
	l := 0
	r := len(a) - 1
	for l < r {
		if a[l] == 1 {
			a[l], a[r] = a[r], a[l]
			r--
		} else {
			l++
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
