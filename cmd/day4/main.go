package main

import (
	"log"
	"os"
	"strings"

	"github.com/xscanpix/adventofcode2024/pkg/day4"
)

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalln("Error reading 'input.txt'. Make sure it exists.")
	}

	text := string(data)

	lineCount := strings.Count(text, "\n")

	textWithoutNewlines := strings.ReplaceAll(text, "\n", "")

	input := day4.Input{
		Data:    &textWithoutNewlines,
		Columns: lineCount + 1,
		Rows:    lineCount + 1,
	}

	log.Printf(`Day3 result 1: %d`, day4.Solve1(input, "XMAS"))
}
