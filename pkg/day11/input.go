package day11

import (
	"container/list"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	StonesList *list.List
	StonesMap  map[int]int
}

type Element struct {
	Value int
}

func ReadInputFromFile(filename string) Input {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	stonesList := list.New()
	stonesMap := make(map[int]int)

	split := strings.Split(string(data), " ")

	for _, s := range split {
		num, _ := strconv.Atoi(s)
		stonesList.PushBack(Element{
			Value: int(num),
		})
		stonesMap[int(num)] = 1
	}

	return Input{
		StonesList: stonesList,
		StonesMap:  stonesMap,
	}
}
