package day12

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

func count_area(input Input, regionId byte, pos pair.Pair) (int, int) {
	current := input.Grid[pos.Y()][pos.X()]

	if current.Visited || current.Id != regionId {
		return 0, 0
	}

	input.Grid[pos.Y()][pos.X()].Visited = true

	adj := nextPairs(pos, input.Width, input.Height)

	areaCount := 1
	perimeter := 4 - len(adj)

	for _, pos := range adj {
		if input.Grid[pos.Y()][pos.X()].Id != regionId {
			perimeter++
		}
		a, p := count_area(input, regionId, pos)
		areaCount += a
		perimeter += p
	}

	return areaCount, perimeter
}

func Solve1(input Input) int {
	sum := 0

	for row := range input.Height {
		for col := range input.Width {
			current := pair.NewPair(col, row)
			if !input.Grid[row][col].Visited {
				a, p := count_area(input, input.Grid[row][col].Id, current)
				sum += a * p
			}
		}
	}

	return sum
}

func Solve2(input Input) int {
	return 0
}
