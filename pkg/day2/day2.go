package day2

import (
	"slices"

	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func Solve(filename string) int {
	arrays := *utils.SplitFileFromFilename(filename, "   ", 2)

	left := arrays[0]
	right := arrays[1]

	slices.Sort(right)

	rightDict := make(map[int]int)

	for i := range right {
		rightDict[right[i]] = rightDict[right[i]] + 1
	}

	sum := 0

	for i := range left {
		l := left[i]

		sum += l * rightDict[l]
	}

	return sum
}
