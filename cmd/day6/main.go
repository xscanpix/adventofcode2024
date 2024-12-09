package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day6"
)

func main() {
	log.Printf(`Day6 result 1: %d`, day6.Solve1(day6.ReadInputFromFile("input.txt")))
	log.Printf(`Day6 result 2: %d`, day6.Solve2(day6.ReadInputFromFile("input.txt")))
}
