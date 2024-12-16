package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/xscanpix/adventofcode2024/internal/pair"
)

type Robot struct {
	P pair.Pair
	V pair.Pair
}

type Input struct {
	Robots []Robot
	Width  int
	Height int
}

func ReadInputFromFile(filename string, biginput bool) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	robots := []Robot{}
	var (
		width, height int
	)

	if biginput {
		width = 101
		height = 103
	} else {
		width = 11
		height = 7
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			px, py, vx, vy int
		)

		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)

		robots = append(robots, Robot{
			P: pair.NewPair(px, py),
			V: pair.NewPair(vx, vy),
		})
	}

	return Input{
		Robots: robots,
		Width:  width,
		Height: height,
	}
}
