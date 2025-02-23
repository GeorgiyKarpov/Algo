package main

import "fmt"

func Solve(target int, a []int) (int, int, error) {
	if len(a) < 2 {
		return 0, 0, fmt.Errorf("n < 2")
	}

	l := 0
	r := len(a) - 1
	for l < r {
		sum := a[l] + a[r]
		if target > sum {
			l++
		} else if target < sum {
			r--
		} else {
			return l, r, nil
		}
	}
	return 0, 0, fmt.Errorf("no sum equals target")
}

func main() {
	var n, target int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Scan(&target)

	ans1, ans2, err := Solve(target, a)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans1, ans2)
	}
}
