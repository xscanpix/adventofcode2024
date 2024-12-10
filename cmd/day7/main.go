package main

import (
	"log"

	"github.com/xscanpix/adventofcode2024/pkg/day7"
)

func main() {
	log.Printf(`Day7 result 1: %d`, day7.Solve1(day7.ReadInputFromFile("input.txt")))
	log.Printf(`Day7 result 2: %d`, day7.Solve2(day7.ReadInputFromFile("input.txt")))
}
