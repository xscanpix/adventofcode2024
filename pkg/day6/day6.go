package day6

import (
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

func nextByte(grid Grid, pos pair.Pair, dir int) byte {
	switch dir {
	case up:
		{
			return grid.Data[pos.Y()-1][pos.X()]
		}
	case down:
		{
			return grid.Data[pos.Y()+1][pos.X()]
		}
	case left:
		{
			return grid.Data[pos.Y()][pos.X()-1]
		}
	case right:
		{
			return grid.Data[pos.Y()][pos.X()+1]
		}
	}

	panic("unreachable")
}

func isNextObstacle(grid Grid, pos pair.Pair, dir int) bool {
	nextByte := nextByte(grid, pos, dir)

	return nextByte == '#'

}

func simulate(input Input) ([]pair.Pair, bool) {
	path := []pair.Pair{}
	obstacles := []pair.Pair{}

	currentPos := pair.NewPair(input.StartX, input.StartY)

	dir := up

	for {
		if input.Grid.Data[currentPos.Y()][currentPos.X()] != 'X' {
			input.Grid.Data[currentPos.Y()][currentPos.X()] = 'X'
			path = append(path, currentPos.Copy())
		}

		if currentPos.X() <= 0 || currentPos.X() >= input.Grid.Width-1 || currentPos.Y() <= 0 || currentPos.Y() >= input.Grid.Height-1 {
			break
		}

		obstacle := false
		if isNextObstacle(input.Grid, currentPos, dir) {
			for i := len(obstacles) - 1; i >= 0; i-- {
				if len(obstacles) > 4 {
					if obstacles[i].X() == currentPos.X() && obstacles[i].Y() == currentPos.Y() {
						return path, true
					}
				}
			}

			obstacle = true
			dir = (dir + 1) % 4
		}

		if obstacle && isNextObstacle(input.Grid, currentPos, dir) {
			for i := len(obstacles) - 1; i >= 0; i-- {
				if len(obstacles) > 4 {
					if obstacles[i].X() == currentPos.X() && obstacles[i].Y() == currentPos.Y() {
						return path, true
					}
				}
			}

			obstacle = true
			dir = (dir + 1) % 4
		}

		if obstacle {
			obstacles = append(obstacles, currentPos.Copy())
		}

		switch dir {
		case up:
			{
				currentPos.DecY()
			}
		case down:
			{
				currentPos.IncY()
			}
		case left:
			{
				currentPos.DecX()
			}
		case right:
			{
				currentPos.IncX()
			}
		}

	}

	return path, false
}

func Solve1(input Input) int {
	path, _ := simulate(input)

	return len(path)
}

func Solve2(input Input) int {
	path, _ := simulate(Input{
		Grid:   input.Grid.Copy(),
		StartX: input.StartX,
		StartY: input.StartY,
	})

	loops := 0

	for _, pos := range path {
		grid := input.Grid.Copy()
		grid.Data[pos.Y()][pos.X()] = '#'

		_, loop := simulate(Input{
			Grid:   grid,
			StartX: input.StartX,
			StartY: input.StartY,
		})

		if loop {
			loops++
		}
	}

	return loops
}
