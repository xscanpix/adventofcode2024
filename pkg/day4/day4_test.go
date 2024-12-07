package day4

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestSolve1(t *testing.T) {
	data, err := os.ReadFile("test_input.txt")

	if err != nil {
		log.Fatalln(err)
	}

	expect := 18

	text := strings.ReplaceAll(string(data), "\n", "")

	input := Input{
		Data:    &text,
		Columns: 10,
		Rows:    10,
	}

	if result := Solve1(input, "XMAS"); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}

func TestSolve2(t *testing.T) {
	data, err := os.ReadFile("test_input2.txt")

	if err != nil {
		log.Fatalln(err)
	}

	expect := 9

	text := strings.ReplaceAll(string(data), "\n", "")

	input := Input{
		Data:    &text,
		Columns: 10,
		Rows:    10,
	}

	if result := Solve2(input, "MAS"); result != expect {
		t.Fatalf(`Expected %d, got %d`, expect, result)
	}
}
