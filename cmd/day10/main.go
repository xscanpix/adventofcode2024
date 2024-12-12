package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day10"
)

func main() {
	log.Printf(`Day10 result 1: %d`, day10.Solve1(day10.ReadInputFromFile("input.txt")))
	log.Printf(`Day10 result 2: %d`, day10.Solve2(day10.ReadInputFromFile("input.txt")))
}
