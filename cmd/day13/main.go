package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day13"
)

func main() {
	log.Printf(`Day13 result 1: %d`, day13.Solve1(day13.ReadInputFromFile("input.txt")))
	log.Printf(`Day13 result 2: %d`, day13.Solve2(day13.ReadInputFromFile("input.txt")))
}
