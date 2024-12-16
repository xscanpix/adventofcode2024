package day14

import (
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func Solve1(input Input) int {
	halfWidth := input.Width / 2
	halfHeight := input.Height / 2

	bounds11 := pair.NewPair(0, 0)
	bounds12 := pair.NewPair(halfWidth, halfHeight)

	bounds21 := pair.NewPair(halfWidth+1, 0)
	bounds22 := pair.NewPair(input.Width, halfHeight)

	bounds31 := pair.NewPair(0, halfHeight+1)
	bounds32 := pair.NewPair(halfWidth, input.Height)

	bounds41 := pair.NewPair(halfWidth+1, halfHeight+1)
	bounds42 := pair.NewPair(input.Width, input.Height)

	quad := func(p pair.Pair) int {
		if p.Bounds(bounds11, bounds12) {
			return 0
		}
		if p.Bounds(bounds21, bounds22) {
			return 1
		}
		if p.Bounds(bounds31, bounds32) {
			return 2
		}
		if p.Bounds(bounds41, bounds42) {
			return 3
		}

		return -1
	}

	quads := [4]int{}
	for _, robot := range input.Robots {
		finalPX := mod(robot.P.X()+(robot.V.X()*100), input.Width)
		finalPY := mod(robot.P.Y()+(robot.V.Y()*100), input.Height)

		finalP := pair.NewPair(int(finalPX), int(finalPY))

		q := quad(finalP)
		if q >= 0 {
			quads[q]++
		}
	}

	return quads[0] * quads[1] * quads[2] * quads[3]
}

func Solve2(input Input) int {
	return 0
}
