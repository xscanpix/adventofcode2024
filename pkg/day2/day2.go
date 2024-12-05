package day2

import (
	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func Solve1(filename string) int {
	reports := *utils.SplitFileIntoRowsFromFilename(filename, " ")

	safe_report_count := 0
	for _, report := range reports {
		report_len := len(report)
		signcount := 0
		safe := true
		for j := 0; j < report_len-1; j++ {
			diff := (report)[j] - (report)[j+1]
			abs := utils.IntAbs(diff)

			if abs < 1 || abs > 3 {
				safe = false
				break
			}

			signcount += (diff / abs)

			if utils.IntAbs(signcount) != j+1 {
				safe = false
				break
			}
		}

		if safe {
			if utils.IntAbs(signcount) == report_len-1 {
				safe_report_count++
			}
		}
	}

	return safe_report_count
}
