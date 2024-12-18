package main

import (
	"log"
	"os"

	"github.com/xscanpix/adventofcode2024/pkg/day3"
)

func main() {
	text, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatalln("Error reading 'input.txt'. Make sure it exists.")
	}

	log.Printf(`Day3 result 1: %d`, day3.Solve1(string(text)))
	log.Printf(`Day3 result 2: %d`, day3.Solve2(string(text)))
}
