package day10

import (
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

func nextPairs(p pair.Pair, width, height int) []pair.Pair {

	res := []pair.Pair{}

	if p.X() > 0 {
		res = append(res, pair.NewPair(p.X()-1, p.Y()))
	}

	if p.X() < width-1 {
		res = append(res, pair.NewPair(p.X()+1, p.Y()))
	}

	if p.Y() > 0 {
		res = append(res, pair.NewPair(p.X(), p.Y()-1))
	}

	if p.Y() < height-1 {
		res = append(res, pair.NewPair(p.X(), p.Y()+1))
	}

	return res
}

func dfsVisit(input Input, vertex pair.Pair, altitude int, trailheads map[int]int) {
	adjacentList := nextPairs(vertex, input.Width, input.Height)

	for _, adjacent := range adjacentList {
		adjacentAltitude := input.Grid[adjacent.Y()][adjacent.X()]
		if (adjacentAltitude + 1) == altitude {
			dfsVisit(input, adjacent, adjacentAltitude, trailheads)

			if adjacentAltitude == 0 {
				trailheads[adjacent.Hash()]++
			}
		} else {
			continue
		}
	}
}

func Solve1(input Input) int {
	trailheadsCount := 0

	for _, start := range input.Start {
		trailheads := make(map[int]int)
		dfsVisit(input, start, 9, trailheads)
		trailheadsCount += len(trailheads)
	}

	return trailheadsCount
}

func Solve2(input Input) int {
	trailheadsCount := 0

	for _, start := range input.Start {
		trailheads := make(map[int]int)
		dfsVisit(input, start, 9, trailheads)

		for _, count := range trailheads {
			trailheadsCount += count
		}
	}

	return trailheadsCount
}
