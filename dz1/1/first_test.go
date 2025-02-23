package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		n        int
		a        []int
		target   int
		expected [2]int
		err      bool
	}{
		{0, []int{}, 2, [2]int{0, 0}, true},                              // n < 2
		{6, []int{1, 2, 3, 4, 5, 6}, 3, [2]int{0, 1}, false},             // 1 + 2 = 3
		{8, []int{3, 8, 9, 11, 16, 18, 19, 21}, 25, [2]int{2, 4}, false}, // 9 + 16 = 25
		{4, []int{3, 8, 9, 1000}, 10000, [2]int{0, 0}, true},             // нет чисел: sum = 10000
	}

	for _, test := range tests {
		ans1, ans2, err := Solve(test.target, test.a)

		if err != nil && test.err != true {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if err == nil && test.err {
			t.Errorf("Expected error, but got none")
		}
		if ans1 != test.expected[0] || ans2 != test.expected[1] {
			t.Errorf("input %v, target %d, expected %v, got (%d, %d)", test.a, test.target, test.expected, ans1, ans2)
		}
	}
}
