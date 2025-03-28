package main

func copyTime(n, x, y int) int {
	var l, r int
	r = max(x, y) * (n - 1)
	mid := 0
	for l+1 < r {
		mid = (l + r) / 2
		if mid/x+mid/y < n-1 {
			l = mid
		} else {
			r = mid
		}
	}

	return min(x, y) + r
}
