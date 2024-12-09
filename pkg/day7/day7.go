package day7

func check(line ParsedLine) int {
	to_check := []ParsedLine{line}

	count := 0

	for len(to_check) > 0 {
		currentLine := to_check[0]
		to_check = to_check[1:]

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

func Solve2(input Input) int {
	return 0
}
