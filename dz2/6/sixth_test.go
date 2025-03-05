package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		a        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
		{"abcba", true},
	}

	for _, tt := range tests {
		ans := Solve(tt.a)
		if ans != tt.expected {
			t.Errorf("input %v, expected %v, but got %v", tt.a, tt.expected, ans)
		}
	}

}
