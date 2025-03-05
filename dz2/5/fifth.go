package main

import "fmt"

func Solve(s1, s2 string) bool {
	a := []rune(s1)
	b := []rune(s2)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		}
		j++
	}

	return i == len(a)
}

func main() {
	var s1, s2 string
	_, err := fmt.Scanln(&s1)
	if err != nil {
		return
	}

	_, err = fmt.Scanln(&s2)
	if err != nil {
		return
	}
	fmt.Println(Solve(s1, s2))
}
