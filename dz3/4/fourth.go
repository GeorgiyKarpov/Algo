package main

func extraLetter(a, b string) rune {
	before := make(map[rune]int)
	for _, ch := range a {
		before[ch]++
	}
	for _, ch := range b {
		if _, ok := before[ch]; ok {
			before[ch]--
			if before[ch] == 0 {
				delete(before, ch)
			}
			continue
		}
		return ch
	}
	return 0
}
