package day8

import (
	"bufio"
	"log"
	"os"

	"github.com/xscanpix/adventofcode2024/internal/pair"
)

type Input struct {
	AntennaMap *map[byte][]pair.Pair
	GridDim    pair.Pair
}

func ReadInputFromFile(filename string) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	scanner := bufio.NewScanner(file)

	antennaMap := make(map[byte][]pair.Pair)

	row := 0
	col := 0
	for scanner.Scan() {
		bytes := scanner.Bytes()
		col = 0
		for _, b := range bytes {
			_, ok := antennaMap[b]

			if b != '.' {
				if !ok {
					antennaMap[b] = make([]pair.Pair, 0)
				}

				antennaMap[b] = append(antennaMap[b], pair.NewPair(row, col))
			}
			col++
		}

		row++
	}

	return Input{
		AntennaMap: &antennaMap,
		GridDim:    pair.NewPair(row, col),
	}
}
