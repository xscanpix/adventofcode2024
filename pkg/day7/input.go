package day7

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type ParsedLine struct {
	Sum     int
	Numbers []int
}

type Input struct {
	Lines []ParsedLine
}

func ReadInputFromFile(filename string) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	scanner := bufio.NewScanner(file)

	parsedLines := []ParsedLine{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		sum, _ := strconv.Atoi(parts[0])
		numbers := strings.Split(parts[1], " ")

		parsedNumbers := make([]int, len(numbers))
		for i := range numbers {
			num, _ := strconv.Atoi(numbers[i])
			parsedNumbers[i] = num
		}

		parsedLine := ParsedLine{
			Sum:     sum,
			Numbers: parsedNumbers,
		}

		parsedLines = append(parsedLines, parsedLine)
	}

	return Input{
		Lines: parsedLines,
	}
}
