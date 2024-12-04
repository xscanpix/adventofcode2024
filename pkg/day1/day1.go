package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xscanpix/adventofcode2024/internal/utils"
)

func Solve(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftSum int
	var rightSum int

	for scanner.Scan() {
		var values = strings.Split(scanner.Text(), "   ")

		left, err := strconv.Atoi(values[0])

		if err != nil {
			log.Fatal(err)
		}

		leftSum += left

		right, err := strconv.Atoi(values[1])

		if err != nil {
			log.Fatal(err)
		}

		rightSum += right
	}

	return utils.IntAbs(leftSum - rightSum)
}
