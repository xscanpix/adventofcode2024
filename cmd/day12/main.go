package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day12"
)

func main() {
	log.Printf(`Day12 result 1: %d`, day12.Solve1(day12.ReadInputFromFile("input.txt")))
	log.Printf(`Day12 result 2: %d`, day12.Solve2(day12.ReadInputFromFile("input.txt")))
}
