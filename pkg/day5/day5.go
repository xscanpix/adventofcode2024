package day5

func Solve1(input Input) int {
	safe_updates := make([]bool, len(input.Updates))
	for i := range safe_updates {
		safe_updates[i] = true
	}

	sum := 0

	for updateIndex, update := range input.Updates {
		seen_before := map[int]bool{}

		for i, number := range update {
			rules := input.Rules[number]
			seen_before[number] = true

			for _, rule := range rules {
				if seen_before[rule] {
					safe_updates[updateIndex] = false
					break
				}
			}

			if !safe_updates[updateIndex] {
				break
			}

			if i == len(update)-1 {
				sum += update[len(update)/2]
			}
		}
	}

	return sum
}

func Solve2(input Input) int {
	return 0
}
