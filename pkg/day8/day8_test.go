package day8

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	expect := 14

	if result := Solve1(ReadInputFromFile("test_input.txt")); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
	expect := 11387

	if result := Solve2(ReadInputFromFile("test_input.txt")); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
