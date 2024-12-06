package day2

import (
	"testing"

	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func TestSolve1(t *testing.T) {
	reports := utils.SplitFileIntoRowsFromFilename("test_input_1.txt", " ")

	expect := 2

	if result := Solve1(*reports); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
	reports := utils.SplitFileIntoRowsFromFilename("test_input_2.txt", " ")

	expect := 4

	if result := Solve2(*reports); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
