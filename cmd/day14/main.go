package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day14"
)

func main() {
	log.Printf(`Day14 result 1: %d`, day14.Solve1(day14.ReadInputFromFile("input.txt", true)))
	log.Printf(`Day14 result 2: %d`, day14.Solve2(day14.ReadInputFromFile("input.txt", true)))
}
