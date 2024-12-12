package day10

import (
	"bufio"
	"log"
	"os"

	"github.com/xscanpix/adventofcode2024/internal/pair"
)

type Input struct {
	Grid   [][]int
	Width  int
	Height int
	Start  []pair.Pair
}

func ReadInputFromFile(filename string) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	start := []pair.Pair{}
	grid := make([][]int, 0)

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		parsed := make([]int, len(line))

		for i := range line {
			num := int(line[i] - 48)
			parsed[i] = num

			if num == 9 {
				start = append(start, pair.NewPair(i, row))
			}
		}

		grid = append(grid, parsed)
		row++
	}

	return Input{
		Grid:   grid,
		Width:  len(grid[0]),
		Height: len(grid),
		Start:  start,
	}
}
