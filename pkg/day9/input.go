package day9

import (
	"log"
	"os"
)

type Input struct {
	Memory        []int
	CompactedSize int
}

func ReadInputFromFile(filename string) Input {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	memory := make([]int, 0)

	compactedSize := 0

	id := 0
	for i := 0; i < len(data); i += 2 {
		blockLen := int(data[i] - 48)
		bSlice := make([]int, blockLen)

		for bi := range bSlice {
			bSlice[bi] = id
			compactedSize++
		}

		memory = append(memory, bSlice...)

		if i+1 < len(data) {
			freespaceLen := int(data[i+1] - 48)

			if freespaceLen > 0 {
				fSlice := make([]int, freespaceLen)

				for f := range fSlice {
					fSlice[f] = -1
				}

				memory = append(memory, fSlice...)
			}
		}

		id++
	}

	return Input{
		Memory:        memory,
		CompactedSize: compactedSize,
	}
}
