package day9

func findNextFreespace(memory []int, prev int) int {
	for i, val := range memory[prev+1:] {
		if val == -1 {
			return prev + (i + 1)
		}
	}

	return -1
}

func compact(input Input) []int {
	memorySize := len(input.Memory)

	nextFreespaceIndex := 0
	for i := len(input.Memory) - 1; i >= 0; i-- {
		if memorySize == input.CompactedSize {
			break
		}
		val := input.Memory[i]

		if val >= 0 {
			nextFreespaceIndex = findNextFreespace(input.Memory, nextFreespaceIndex)
			input.Memory[nextFreespaceIndex] = val
			input.Memory[i] = -1
		}

		memorySize--
	}

	return input.Memory[:memorySize]
}

func Solve1(input Input) int {
	compacted := compact(input)

	sum := 0

	for i, val := range compacted {
		sum += i * val
	}

	return sum
}

func Solve2(input Input) int {
	return 0
}
