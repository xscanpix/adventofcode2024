package day11

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	expect := 55312

	if result := Solve1(ReadInputFromFile("test_input.txt")); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
