package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day15"
)

func main() {
	log.Printf(`Day15 result 1: %d`, day15.Solve1(day15.ReadInputFromFile("input.txt", true)))
	log.Printf(`Day15 result 2: %d`, day15.Solve2(day15.ReadInputFromFile("input.txt", true)))
}
