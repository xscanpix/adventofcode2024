package day1

import (
	"slices"
)

func Solve1(left []int, right []int) int {
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

func Solve2(left []int, right []int) int {
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
