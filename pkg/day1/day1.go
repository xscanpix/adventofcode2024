package day1

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(filename string) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left []int
	var right []int

	for scanner.Scan() {
		var values = strings.Split(scanner.Text(), "   ")

		leftVal, err := strconv.Atoi(values[0])

		if err != nil {
			log.Fatal(err)
		}

		left = append(left, leftVal)

		rightVal, err := strconv.Atoi(values[1])

		if err != nil {
			log.Fatal(err)
		}

		right = append(right, rightVal)
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		l := left[i]
		r := right[i]

		if l <= r {
			left[i] = r - l
		} else {
			left[i] = l - r
		}
	}

	sum := 0

	for i := range left {
		sum += left[i]
	}

	return sum
}
