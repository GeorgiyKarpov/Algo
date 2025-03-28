package main

func binarySearchSqrt(target int) int {
	var l, r int
	r = target
	middle := 0
	for l <= r {
		middle = (l + r) / 2
		if middle*middle > target {
			r = middle - 1
		} else if middle*middle < target {
			l = middle + 1
		} else {
			return middle
		}
	}

	if (r+1)*(r+1)-target < target-r*r {
		return r + 1
	}
	return r
}
