package day3

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func do_calc(input string) int {
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

func Solve1(input string) int {
	return do_calc(input)
}

func Solve2(input string) int {
	enabled := true
	sum := 0

	for {
		if enabled {
			index := strings.Index(input, "don't()")

			if index < 0 {
				sum += do_calc(input)
				break
			} else {
				sum += do_calc(input[:index])
			}

			input = input[index+7:]
			enabled = false
		} else {
			index := strings.Index(input, "do()")

			if index < 0 {
				break
			}

			input = input[index+4:]
			enabled = true
		}
	}

	return sum
}
