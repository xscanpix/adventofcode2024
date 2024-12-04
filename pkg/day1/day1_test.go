package day1

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	filename := "test_input_1.txt"
	expect := 11

	if result := Solve1(filename); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
	filename := "test_input_1.txt"
	expect := 31

	if result := Solve2(filename); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
