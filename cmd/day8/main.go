package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day8"
)

func main() {
	log.Printf(`Day8 result 1: %d`, day8.Solve1(day8.ReadInputFromFile("input.txt")))
	log.Printf(`Day8 result 2: %d`, day8.Solve2(day8.ReadInputFromFile("input.txt")))
}
