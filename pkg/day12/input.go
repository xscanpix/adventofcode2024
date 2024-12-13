package day12

import (
	"bufio"
	"log"
	"os"
)

type Plot struct {
	Id      byte
	Visited bool
}

type Input struct {
	Grid   [][]Plot
	Width  int
	Height int
}

func ReadInputFromFile(filename string) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	grid := make([][]Plot, 0)

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		plots := make([]Plot, len(line))

		for i := range line {
			plots[i] = Plot{
				Id:      line[i],
				Visited: false,
			}
		}

		grid = append(grid, plots)

		row++
	}

	return Input{
		Grid:   grid,
		Width:  len(grid[0]),
		Height: len(grid),
	}
}
