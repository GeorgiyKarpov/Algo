package main

import "testing"

func TestExtraLetter(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected rune
	}{
		{
			a:        "",
			b:        "z",
			expected: 'z',
		},
		{
			a:        "uio",
			b:        "oeiu",
			expected: 'e',
		},
		{
			a:        "fe",
			b:        "efo",
			expected: 'o',
		},
		{
			a:        "ab",
			b:        "ab",
			expected: 0,
		},
		{
			a:        "bbb",
			b:        "bbbb",
			expected: 'b',
		},
	}

	for _, test := range tests {
		ans := extraLetter(test.a, test.b)
		if ans != test.expected {
			t.Errorf("input %v %v, expected %v, but got %v", test.a, test.b, test.expected, ans)
		}
	}
}
