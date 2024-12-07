package day5

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Rules   map[int][]int
	Updates [][]int
}

const (
	order  = 0
	update = 1
)

func parseRules(rules []string) map[int][]int {
	parsed := map[int][]int{}

	for _, rule := range rules {
		parts := strings.Split(rule, "|")

		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])

		parsed[left] = append(parsed[left], right)
	}

	return parsed
}

func parseUpdates(updates []string) [][]int {
	parsedUpdates := make([][]int, len(updates))

	for i, update := range updates {
		parts := strings.Split(update, ",")

		parsedUpdate := make([]int, len(parts))

		for j, part := range parts {
			val, _ := strconv.Atoi(part)
			parsedUpdate[j] = val
		}

		parsedUpdates[i] = parsedUpdate
	}

	return parsedUpdates
}

func ReadInputFromFile(filename string) Input {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	state := order

	pageOrderRules := []string{}
	updates := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			state = update
			continue
		}

		switch state {
		case order:
			{
				pageOrderRules = append(pageOrderRules, line)
			}
		case update:
			{
				updates = append(updates, line)
			}
		}
	}

	return Input{
		Rules:   parseRules(pageOrderRules),
		Updates: parseUpdates(updates),
	}
}
