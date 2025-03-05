package main

import (
	"fmt"
	"unicode"
)

func Solve(s1 string) bool {
	var a []rune
	for _, i := range s1 {
		if unicode.IsLetter(i) || unicode.IsDigit(i) {
			a = append(a, unicode.ToLower(i))
		}
	}
	left := 0
	right := len(a) - 1
	for left < right {
		if a[left] != a[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	var s1 string
	_, err := fmt.Scanln(&s1)
	if err != nil {
		return
	}
	Solve(s1)
}
