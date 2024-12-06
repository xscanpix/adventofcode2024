package day2

import (
	"github.com/xscanpix/adventofcode2024/internal/utils"
)

type pair struct {
	indexA int
	indexB int
}

func is_level_change_safe(report *[]int, pair *pair) bool {
	abs_diff := utils.IntAbs((*report)[pair.indexA] - (*report)[pair.indexB])

	if abs_diff >= 1 && abs_diff <= 3 {
		return true
	}

	return false
}

func compare_level_pairs(report *[]int, pair1 *pair, pair2 *pair) bool {
	diff_1 := (*report)[pair1.indexA] - (*report)[pair1.indexB]
	diff_2 := (*report)[pair2.indexA] - (*report)[pair2.indexB]

	if (diff_1 < 0 && diff_2 < 0) || (diff_1 > 0 && diff_2 > 0) {
		if is_level_change_safe(report, pair1) && is_level_change_safe(report, pair2) {
			return true
		}
	}

	return false
}

func Solve1(filename string) int {
	reports := *utils.SplitFileIntoRowsFromFilename(filename, " ")

	safe_report_count := 0
	for _, report := range reports {
		// Assume safe
		safe_report_count += 1
		for i := 1; i < len(report)-1; i++ {
			if !compare_level_pairs(&report, &pair{indexA: i - 1, indexB: i}, &pair{indexA: i, indexB: i + 1}) {
				// Retract the safe increase
				safe_report_count -= 1
				break
			}
		}
	}

	return safe_report_count
}

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func find_unsafe_index(levels []int) int {
	for i := range len(levels) - 2 {
		pair1 := &pair{indexA: i, indexB: i + 1}
		pair2 := &pair{indexA: i + 1, indexB: i + 2}

		diff_1 := levels[pair1.indexA] - levels[pair1.indexB]
		diff_2 := levels[pair2.indexA] - levels[pair2.indexB]

		pair1_level_change_safe := is_level_change_safe(&levels, pair1)
		pair2_level_change_safe := is_level_change_safe(&levels, pair2)

		if (diff_1 < 0 && diff_2 < 0) || (diff_1 > 0 && diff_2 > 0) {
			if !pair1_level_change_safe || !pair2_level_change_safe {
				return i
			}
		} else if diff_1 == 0 || diff_2 == 0 {
			return i
		} else {
			return i
		}
	}

	return -1
}

func Solve2(filename string) int {
	reports := *utils.SplitFileIntoRowsFromFilename(filename, " ")

	safe_report_count := 0
	for _, report := range reports {
		unsafe_index := find_unsafe_index(report)
		if unsafe_index >= 0 {

			for i := unsafe_index; i < unsafe_index+3; i++ {
				report_copy := make([]int, len(report))
				copy(report_copy, report)

				report_copy = removeIndex(report_copy, i)
				checked := find_unsafe_index(report_copy)
				if checked < 0 {
					safe_report_count++
					break
				}
			}
		} else {
			safe_report_count++
		}
	}

	return safe_report_count
}
