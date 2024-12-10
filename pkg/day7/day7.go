package day7

import (
	"math"
)

func check(line ParsedLine) int {
	to_check := []ParsedLine{line}

	count := 0

	for len(to_check) > 0 {
		currentLine := to_check[0]
		to_check = to_check[1:]

		if len(currentLine.Numbers) <= 1 {
			continue
		}

		currentSum := currentLine.Sum
		for i := len(currentLine.Numbers) - 1; i > 0; i-- {
			currentNum := currentLine.Numbers[i]

			if (currentSum % currentNum) == 0 {
				to_check = append(to_check, ParsedLine{
					Sum:     currentSum - currentNum,
					Numbers: currentLine.Numbers[:i],
				})
				currentSum = currentSum / currentNum
			} else {
				currentSum = currentSum - currentNum
			}
		}

		if currentSum == currentLine.Numbers[0] {
			count++
		}
	}

	return count
}

func Solve1(input Input) int {
	sum := 0

	for _, line := range input.Lines {
		if check(line) > 0 {
			sum += line.Sum
		}
	}

	return sum
}

func recurse(sum int, numbers []int) bool {

	if len(numbers) == 0 {
		panic("Impossible condition")
	}

	if len(numbers) == 1 {
		return sum == numbers[0]
	}

	num := numbers[len(numbers)-1]

	add := ((sum >= num) && recurse(sum-num, numbers[:len(numbers)-1]))
	mul := ((sum%num) == 0 && recurse(sum/num, numbers[:len(numbers)-1]))

	p := int(math.Pow10(int(math.Log10(float64(num)) + 1)))

	conc := (((sum % p) == num) && recurse(sum/p, numbers[:len(numbers)-1]))

	return add || mul || conc
}

func Solve2(input Input) int {
	sum := 0

	for _, line := range input.Lines {
		if recurse(line.Sum, line.Numbers) {
			sum += line.Sum
		}
	}

	return sum
}
