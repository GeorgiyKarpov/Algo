package main

import "fmt"

func Solve(n uint64) uint64 {
	var sum uint64
	sum = n * (n + 1) / 2
	return sum
}

func main() {
	var n uint64
	fmt.Scan(&n)
	fmt.Println(Solve(n))
}
