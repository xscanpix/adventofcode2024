package day11

import (
	"container/list"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	Numbers *list.List
}

type Element struct {
	Value int
}

func ReadInputFromFile(filename string) Input {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	numbers := list.New()
	split := strings.Split(string(data), " ")

	for _, s := range split {
		num, _ := strconv.Atoi(s)
		numbers.PushBack(Element{
			Value: num,
		})
	}

	return Input{
		Numbers: numbers,
	}
}
