package day8

import (
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

func find_antinodes(antenna1 pair.Pair, antenna2 pair.Pair, distance int, gridDim pair.Pair) []pair.Pair {
	nodes := make([]pair.Pair, 0)

	xDiff := antenna1.X() - antenna2.X()
	yDiff := antenna1.Y() - antenna2.Y()

	antinode1 := pair.NewPair(antenna1.X()+(xDiff*distance), antenna1.Y()+(yDiff*distance))
	antinode2 := pair.NewPair(antenna2.X()-(xDiff*distance), antenna2.Y()-(yDiff*distance))

	if antinode1.Bounds(pair.Zero(), gridDim) {
		nodes = append(nodes, antinode1)
	}

	if antinode2.Bounds(pair.Zero(), gridDim) {
		nodes = append(nodes, antinode2)
	}

	return nodes
}

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

				for _, node := range find_antinodes(antenna1, antenna2, 1, input.GridDim) {
					unique_antinodes[node.Hash()] = true
				}
			}
		}
	}

	return len(unique_antinodes)
}

func Solve2(input Input) int {
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

				distance := 0
				for {
					antinodes := find_antinodes(antenna1, antenna2, distance, input.GridDim)

					if len(antinodes) == 0 {
						break
					}

					for _, node := range antinodes {
						unique_antinodes[node.Hash()] = true
					}

					distance++
				}
			}
		}
	}

	return len(unique_antinodes)
}
