package day2

import (
	"testing"
)

func TestSolve1(t *testing.T) {
	filename := "test_input_1.txt"
	expect := 31

	if result := Solve(filename); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
