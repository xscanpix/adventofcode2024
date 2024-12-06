package day2

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	filename := "test_input_1.txt"
	expect := 2

	if result := Solve1(filename); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
	filename := "test_input_2.txt"
	expect := 4

	if result := Solve2(filename); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
