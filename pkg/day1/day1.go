package day1

import (
	"slices"

	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func Solve1(filename string) int {
	arrays := utils.SplitFileFromFilename(filename, "   ", 2)

	left := &(*arrays)[0]
	right := &(*arrays)[1]

	slices.Sort(*left)
	slices.Sort(*right)

	for i := range *left {
		l := (*left)[i]
		r := (*right)[i]

		if l <= r {
			(*left)[i] = r - l
		} else {
			(*left)[i] = l - r
		}
	}

	sum := 0

	for i := range *left {
		sum += (*left)[i]
	}

	return sum
}

func Solve2(filename string) int {
	// Setup code
	arrays := utils.SplitFileFromFilename(filename, "   ", 2)

	left := &(*arrays)[0]
	right := &(*arrays)[1]

	// Create map to store unique sums
	rightDict := make(map[int]int)

	// Populate map
	for i := range *right {
		rightDict[(*right)[i]] = rightDict[(*right)[i]] + 1
	}

	// Perform summation between left and right
	sum := 0
	for i := range *left {
		l := (*left)[i]

		sum += l * rightDict[l]
	}

	return sum
}
