package day2

import (
	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func Solve(filename string) int {
	// Setup code
	arrays := *utils.SplitFileFromFilename(filename, "   ", 2)

	left := arrays[0]
	right := arrays[1]

	// Create map to store unique sums
	rightDict := make(map[int]int)

	// Populate map
	for i := range right {
		rightDict[right[i]] = rightDict[right[i]] + 1
	}

	// Perform summation between left and right
	sum := 0
	for i := range left {
		l := left[i]

		sum += l * rightDict[l]
	}

	return sum
}
