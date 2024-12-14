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

func part2(input Input, regionId byte, pos pair.Pair) (int, int) {
	current := input.Grid[pos.Y()][pos.X()]

	if current.Visited || current.Id != regionId {
		return 0, 0
	}

	input.Grid[pos.Y()][pos.X()].Visited = true

	adj := nextPairs(pos, input.Width, input.Height)

	corners := 0
	areaCount := 1
	perimeter := 4 - len(adj)

	left := byte(0)
	right := byte(0)
	up := byte(0)
	down := byte(0)

	upleft := byte(0)
	upright := byte(0)
	downleft := byte(0)
	downright := byte(0)

	if pos.X() > 0 {
		left = input.Grid[pos.Y()][pos.X()-1].Id
	}

	if pos.X() < input.Width-1 {
		right = input.Grid[pos.Y()][pos.X()+1].Id
	}

	if pos.Y() > 0 {
		up = input.Grid[pos.Y()-1][pos.X()].Id
		if pos.X() > 0 {
			upleft = input.Grid[pos.Y()-1][pos.X()-1].Id
		}

		if pos.X() < input.Width-1 {
			upright = input.Grid[pos.Y()-1][pos.X()+1].Id
		}
	}

	if pos.Y() < input.Height-1 {
		down = input.Grid[pos.Y()+1][pos.X()].Id
		if pos.X() > 0 {
			downleft = input.Grid[pos.Y()+1][pos.X()-1].Id
		}

		if pos.X() < input.Width-1 {
			downright = input.Grid[pos.Y()+1][pos.X()+1].Id
		}
	}

	// Outside corners
	if left != regionId && up != regionId {
		corners++
	}

	if left != regionId && down != regionId {
		corners++
	}

	if up != regionId && right != regionId {
		corners++
	}

	if right != regionId && down != regionId {
		corners++
	}

	// Inside corners
	if right == regionId && down == regionId && downright != regionId {
		corners++
	}

	if right == regionId && up == regionId && upright != regionId {
		corners++
	}

	if left == regionId && up == regionId && upleft != regionId {
		corners++
	}

	if left == regionId && down == regionId && downleft != regionId {
		corners++
	}

	for _, pos := range adj {
		if input.Grid[pos.Y()][pos.X()].Id != regionId {
			perimeter++
		}
		a, c := part2(input, regionId, pos)
		corners += c
		areaCount += a
	}

	return areaCount, corners
}

func Solve2(input Input) int {
	sum := 0

	for row := range input.Height {
		for col := range input.Width {
			current := pair.NewPair(col, row)
			if !input.Grid[row][col].Visited {
				a, c := part2(input, input.Grid[row][col].Id, current)
				sum += a * c
			}
		}
	}

	return sum
}
