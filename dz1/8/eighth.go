package main

import "fmt"

func Solve(a []int) []int {
	even := 0
	l := 0
	for l < len(a) {
		if a[l]%2 == 0 {
			a[l], a[even] = a[even], a[l]
			even++
		}
		l++
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
