package day1

import (
	"testing"

	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func TestSolve1(t *testing.T) {
	arrays := utils.SplitFileIntoColumnsFromFilename("test_input_1.txt", "   ", 2)

	expect := 11

	if result := Solve1((*arrays)[0], (*arrays)[1]); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
	arrays := utils.SplitFileIntoColumnsFromFilename("test_input_1.txt", "   ", 2)

	expect := 31

	if result := Solve2((*arrays)[0], (*arrays)[1]); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
