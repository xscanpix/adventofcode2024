package day15

import (
	"bufio"
	"log"
	"os"

	"github.com/xscanpix/adventofcode2024/internal/grid"
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

const (
	Wall  = byte('#')
	Empty = byte('.')
	Box   = byte('O')
	Robot = byte('@')
)

const (
	Up    = byte('^')
	Right = byte('>')
	Down  = byte('v')
	Left  = byte('<')
)

type Input struct {
	Map   grid.Grid[byte]
	Start pair.Pair
	Moves []byte
}

func ReadInputFromFile(filename string, biginput bool) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	scanner := bufio.NewScanner(file)

	data := []byte{}

	var start pair.Pair

	rows := 0
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			break
		}

		for i, m := range line {
			if m == '#' {
				data = append(data, Wall)
			} else if m == '.' {
				data = append(data, Empty)
			} else if m == 'O' {
				data = append(data, Box)
			} else {
				data = append(data, Robot)
				start = pair.NewPair(i, rows)
			}
		}

		rows++
	}

	moves := []byte{}

	for scanner.Scan() {
		line := scanner.Bytes()

		for _, m := range line {
			if m == '^' {
				moves = append(moves, Up)
			} else if m == '>' {
				moves = append(moves, Right)
			} else if m == 'v' {
				moves = append(moves, Down)
			} else {
				moves = append(moves, Left)
			}
		}
	}

	return Input{
		Map:   grid.NewGrid(data, len(data)/rows, rows),
		Start: start,
		Moves: moves,
	}
}
