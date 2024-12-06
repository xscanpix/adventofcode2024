package main

import (
	"log"
	"slices"

	"github.com/xscanpix/adventofcode2024/internal/utils"
	"github.com/xscanpix/adventofcode2024/pkg/day2"
)

func main() {
	arrays := utils.SplitFileIntoRowsFromFilename("input.txt", " ")

	log.Printf(`Day2 result 1: %d`, day2.Solve1(slices.Clone(*arrays)))
	log.Printf(`Day2 result 2: %d`, day2.Solve2(slices.Clone(*arrays)))
}
