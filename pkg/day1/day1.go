package day1

import (
	"slices"

	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func Solve(filename string) int {
	arrays := *utils.SplitFileFromFilename(filename, "   ", 2)

	left := arrays[0]
	right := arrays[1]

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		l := left[i]
		r := right[i]

		if l <= r {
			left[i] = r - l
		} else {
			left[i] = l - r
		}
	}

	sum := 0

	for i := range left {
		sum += left[i]
	}

	return sum
}
