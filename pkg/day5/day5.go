package day5

func findErrorIndexes(update []int, rulesMap map[int][]int) []int {
	seen_before := map[int]int{}

	for numberIndex, number := range update {
		rules := rulesMap[number]
		seen_before[number] = numberIndex

		for _, rule := range rules {
			val, ok := seen_before[rule]

			if ok {
				return []int{val, numberIndex}
			}
		}
	}

	return nil
}

func Solve1(input Input) int {
	sum := 0

	for _, update := range input.Updates {
		errorIndexes := findErrorIndexes(update, input.Rules)

		if errorIndexes == nil {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func Solve2(input Input) int {
	sum := 0

	for _, update := range input.Updates {
		errorIndexes := findErrorIndexes(update, input.Rules)

		if errorIndexes != nil {
			for {
				temp := update[errorIndexes[0]]
				update[errorIndexes[0]] = update[errorIndexes[1]]
				update[errorIndexes[1]] = temp

				errorIndexes = findErrorIndexes(update, input.Rules)

				if errorIndexes == nil {
					break
				}
			}

			sum += update[len(update)/2]
		}
	}

	return sum
}
