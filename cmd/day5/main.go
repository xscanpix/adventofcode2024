package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day5"
)

func main() {
	log.Printf(`Day5 result 1: %d`, day5.Solve1(day5.ReadInputFromFile("input.txt")))
	log.Printf(`Day5 result 2: %d`, day5.Solve2(day5.ReadInputFromFile("input.txt")))
}
