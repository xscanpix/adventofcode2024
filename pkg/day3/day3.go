package day3

import (
	"log"
	"regexp"
	"strconv"
)

func Solve1(input string) int {
	reg, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)

	if err != nil {
		log.Fatal(err)
	}

	sum := 0

	for {
		index := reg.FindStringSubmatchIndex(input)

		if index == nil {
			break
		}

		val1, err1 := strconv.Atoi(string(input[index[2]:index[3]]))
		val2, err2 := strconv.Atoi(string(input[index[4]:index[5]]))

		if err != nil && err2 != nil {
			log.Fatalf(`%v, %v`, err1, err2)
		}

		mul := val1 * val2

		sum = sum + mul

		input = input[index[1]:]
	}

	return sum
}
