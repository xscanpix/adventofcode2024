package day9

import (
	"slices"
)

func findNextFreespaceStart(memory []int, start int) int {
	for i, val := range memory[start:] {
		if val == -1 {
			return start + i
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
			nextFreespaceIndex = findNextFreespaceStart(input.Memory, nextFreespaceIndex)
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
	left := 0

	for left < len(input.Sections) {
		if input.Sections[left].Value == -1 {
			free := &input.Sections[left]
			right := len(input.Sections) - 1
			for right > left {
				if input.Sections[right].Value >= 0 {
					file := input.Sections[right]
					if file.Size <= free.Size {
						free.Size -= file.Size
						input.Sections = slices.Concat(input.Sections[0:right], []Section{{Size: file.Size, Value: -1}}, input.Sections[right+1:len(input.Sections)])
						input.Sections = slices.Concat(input.Sections[0:left], []Section{{Size: file.Size, Value: file.Value}}, input.Sections[left:len(input.Sections)])
						break
					}
				}
				right--
			}
		}
		left++
	}

	sum := 0

	index := 0
	for _, section := range input.Sections {
		if section.Size > 0 {
			for range section.Size {
				if section.Value >= 0 {
					sum += section.Value * index
				}
				index++
			}
		}
	}

	return sum
}
