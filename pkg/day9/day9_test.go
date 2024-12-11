package day9

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	expect := 1928

	if result := Solve1(ReadInputFromFile("test_input.txt")); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
}
