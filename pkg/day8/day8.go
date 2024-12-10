package day8

import (
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

func Solve1(input Input) int {
	unique_antinodes := make(map[int]bool)

	for _, v := range *input.AntennaMap {
		if len(v) < 2 {
			continue
		}

		for i := range len(v) {
			for j := range len(v) {
				if i == j {
					continue
				}

				antenna1 := v[i]
				antenna2 := v[j]

				xDiff := antenna1.X() - antenna2.X()
				yDiff := antenna1.Y() - antenna2.Y()

				antinode1 := pair.NewPair(antenna1.X()+xDiff, antenna1.Y()+yDiff)
				antinode2 := pair.NewPair(antenna2.X()-xDiff, antenna2.Y()-yDiff)

				if antinode1.Bounds(pair.Zero(), input.GridDim) {
					unique_antinodes[antinode1.Hash()] = true
				}

				if antinode2.Bounds(pair.Zero(), input.GridDim) {
					unique_antinodes[antinode2.Hash()] = true
				}
			}
		}
	}

	return len(unique_antinodes)
}

func Solve2(input Input) int {
	return 0
}
