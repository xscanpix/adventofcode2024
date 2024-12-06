package main

import (
	"log"
	"slices"

	"github.com/xscanpix/adventofcode2024/internal/utils"
	"github.com/xscanpix/adventofcode2024/pkg/day1"
)

func main() {
	arrays := utils.SplitFileIntoColumnsFromFilename("input.txt", "   ", 2)

	log.Printf(`Day1 result 1: %d`, day1.Solve1(slices.Clone((*arrays)[0]), slices.Clone((*arrays)[1])))
	log.Printf(`Day1 result 2: %d`, day1.Solve2(slices.Clone((*arrays)[0]), slices.Clone((*arrays)[1])))
}
