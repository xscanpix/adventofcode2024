package day15

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	expect := 10092

	if result := Solve1(ReadInputFromFile("test_input.txt", false)); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
	expect := 480

	if result := Solve2(ReadInputFromFile("test_input.txt", true)); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
