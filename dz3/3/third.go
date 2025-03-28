package main

import "sort"

func feedAnimals(animals, food []int) int {
	if len(animals) == 0 || len(food) == 0 {
		return 0
	}
	sort.Ints(animals)
	sort.Ints(food)

	l := 0
	r := 0
	ans := 0
	for l < len(animals) && r < len(food) {
		if food[r] >= animals[l] {
			ans++
			r++
			l++
		} else {
			r++
		}
	}
	return ans
}
