package main

import "fmt"

func reverseArr(a []int, l, r int) []int {
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return a
}

func Solve(a []int, n, k int) []int {
	if n == 0 {
		return a
	}
	a = reverseArr(a, 0, n-1)
	a = reverseArr(a, 0, k%n-1)
	a = reverseArr(a, k%n, n-1)
	return a
}

func main() {
	var n, k int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)

	fmt.Println(Solve(a, n, k))
}
