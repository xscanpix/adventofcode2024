package day15

import (
	"errors"
	"log"

	"github.com/xscanpix/adventofcode2024/internal/grid"
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

func printMap(grid grid.Grid[byte]) {
	for j := range grid.Height {
		log.Println(string(grid.Data[j*grid.Width : (j*grid.Width)+grid.Width]))
	}
}

func nextPosition(move byte, currentPos pair.Pair, width, height int) (pair.Pair, error) {
	if move == Up {
		if currentPos.Y() > 0 {
			return pair.NewPair(currentPos.X(), currentPos.Y()-1), nil
		} else {
			return currentPos, errors.New("out-of-bounds")
		}
	} else if move == Right {
		if currentPos.X() < width-1 {
			return pair.NewPair(currentPos.X()+1, currentPos.Y()), nil
		} else {
			return currentPos, errors.New("out-of-bounds")
		}
	} else if move == Down {
		if currentPos.Y() < height-1 {
			return pair.NewPair(currentPos.X(), currentPos.Y()+1), nil
		} else {
			return currentPos, errors.New("out-of-bounds")
		}
	} else {
		// Left
		if currentPos.X() > 0 {
			return pair.NewPair(currentPos.X()-1, currentPos.Y()), nil
		} else {
			return currentPos, errors.New("out-of-bounds")
		}
	}
}

func countBoxesAndSpaces(dir byte, pos pair.Pair, grid grid.Grid[byte]) (int, int) {
	boxes := 0

	for {
		nextPos, err := nextPosition(dir, pos, grid.Width, grid.Height)
		if err != nil {
			panic(err)
		}
		nextByte := grid.GetByte(nextPos.X(), nextPos.Y())

		if nextByte == Box {
			boxes++
		} else {
			break
		}

		pos = nextPos
	}

	for {
		nextPos, err := nextPosition(dir, pos, grid.Width, grid.Height)
		if err != nil {
			panic(err)
		}

		nextByte := grid.GetByte(nextPos.X(), nextPos.Y())

		if nextByte == Empty {
			return boxes, 1
		} else {
			return boxes, 0
		}
	}
}

func moveRobot(move byte, pos pair.Pair, grid grid.Grid[byte]) pair.Pair {
	nextPos, err := nextPosition(move, pos, grid.Width, grid.Height)

	if err != nil {
		panic(err)
	}

	nextByte := grid.GetByte(nextPos.X(), nextPos.Y())

	if nextByte == Wall {
		return pos.Copy()
	} else if nextByte == Box {
		boxes, spaces := countBoxesAndSpaces(move, pos, grid)

		if spaces > 0 {
			// Move unit
			var from pair.Pair
			var to pair.Pair

			for i := boxes; i >= 0; i-- {
				if move == Up {
					from = pair.NewPair(pos.X(), pos.Y()-i)
					to = pair.NewPair(pos.X(), pos.Y()-spaces-i)
				} else if move == Down {
					from = pair.NewPair(pos.X(), pos.Y()+i)
					to = pair.NewPair(pos.X(), pos.Y()+spaces+i)
				} else if move == Right {
					from = pair.NewPair(pos.X()+i, pos.Y())
					to = pair.NewPair(pos.X()+spaces+i, pos.Y())
				} else {
					// Left
					from = pair.NewPair(pos.X()-i, pos.Y())
					to = pair.NewPair(pos.X()-spaces-i, pos.Y())
				}

				t := grid.GetByte(from.X(), from.Y())
				grid.SetByte(Empty, from.X(), from.Y())
				grid.SetByte(t, to.X(), to.Y())
			}

			return to.Copy()
		} else {
			return pos.Copy()
		}
	} else {
		grid.SetByte(Empty, pos.X(), pos.Y())
		grid.SetByte(Robot, nextPos.X(), nextPos.Y())
		return nextPos.Copy()
	}
}

func gpsCount(grid grid.Grid[byte]) int {
	sum := 0

	for j := range grid.Height {
		for i := range grid.Width {
			if grid.GetByte(i, j) == Box {
				sum += 100*j + i
			}
		}
	}

	return sum
}

func Solve1(input Input) int {
	pos := input.Start

	for _, move := range input.Moves {
		pos = moveRobot(move, pos, input.Map)
	}

	return gpsCount(input.Map)
}

func Solve2(input Input) int {
	return 0
}
