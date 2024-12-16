package day13

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xscanpix/adventofcode2024/internal/pair"
)

type Machine struct {
	Buttons [2]pair.Pair
	Prize   pair.Pair
}

type Input struct {
	Machines []Machine
}

func ReadInputFromFile(filename string) Input {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf(`Error reading '%s'. Make sure it exists.`, filename)
	}

	lines := strings.Split(string(data), "\n")

	machines := make([]Machine, 0)

	for i := 0; i < len(lines)-1; i += 4 {
		machine := Machine{}

		// Button A
		buttonA := strings.Split(strings.Split(lines[i], ": ")[1], ", ")
		buttonAX, _ := strconv.Atoi(buttonA[0][2:])
		buttonAY, _ := strconv.Atoi(buttonA[1][2:])
		machine.Buttons[0] = pair.NewPair(buttonAX, buttonAY)

		// Button B
		buttonB := strings.Split(strings.Split(lines[i+1], ": ")[1], ", ")
		buttonBX, _ := strconv.Atoi(buttonB[0][2:])
		buttonBY, _ := strconv.Atoi(buttonB[1][2:])
		machine.Buttons[1] = pair.NewPair(buttonBX, buttonBY)

		// Prize
		prize := strings.Split(strings.Split(lines[i+2], ": ")[1], ", ")
		prizeX, _ := strconv.Atoi(prize[0][2:])
		prizeY, _ := strconv.Atoi(prize[1][2:])
		machine.Prize = pair.NewPair(prizeX, prizeY)

		machines = append(machines, machine)
	}

	return Input{
		Machines: machines,
	}
}
