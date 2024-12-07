package day6

import (
	"bufio"
	"log"
	"os"
)

type Grid struct {
	Data   [][]byte
	Width  int
	Height int
}

type Input struct {
	Grid   Grid
	StartX int
	StartY int
}

func ReadInputFromFile(filename string) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	scanner := bufio.NewScanner(file)

	grid := Grid{
		Data: make([][]byte, 0),
	}
	startX := 0
	startY := 0

	rows := 0

	for scanner.Scan() {
		bytes := scanner.Bytes()

		grid.Data = append(grid.Data, make([]byte, len(bytes)))
		for i := range bytes {
			grid.Data[rows][i] = bytes[i]

			if bytes[i] == '^' {
				startX = i
				startY = rows
			}
		}

		rows++
	}

	grid.Width = len(grid.Data[0])
	grid.Height = len(grid.Data)

	return Input{
		Grid:   grid,
		StartX: startX,
		StartY: startY,
	}
}
