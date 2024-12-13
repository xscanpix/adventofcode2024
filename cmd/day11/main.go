package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day11"
)

func main() {
	log.Printf(`Day11 result 1: %d`, day11.Solve1(day11.ReadInputFromFile("input.txt")))
	log.Printf(`Day11 result 2: %d`, day11.Solve2(day11.ReadInputFromFile("input.txt")))
}
