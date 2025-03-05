package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
	}{
		{"abd", "uabqd", true},
		{"a", "a", true},
		{"abc", "", false},
		{"abbbbbb", "a123bbbbbb", true},
		{"", "abc", true},
	}

	for _, tt := range tests {
		ans := Solve(tt.a, tt.b)
		if ans != tt.expected {
			t.Errorf("input %v, %v, expected %v, but got %v", tt.a, tt.b, tt.expected, ans)
		}
	}

}
