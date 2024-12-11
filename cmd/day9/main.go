package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day9"
)

func main() {
	log.Printf(`Day9 result 1: %d`, day9.Solve1(day9.ReadInputFromFile("input.txt")))
	log.Printf(`Day9 result 2: %d`, day9.Solve2(day9.ReadInputFromFile("input.txt")))
}
