package day14

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	expect := 12

	if result := Solve1(ReadInputFromFile("test_input.txt", false)); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
