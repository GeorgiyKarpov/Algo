package main

import "testing"

func TestZero(t *testing.T) {
	sum := Solve(100)
	var expected uint64
	expected = 5050

	if sum != expected {
		t.Errorf("expected %d, but got %d", expected, sum)
	}
}
