package day11

import (
	"container/list"
	"math"
)

func blinkMap(stones map[int]int) {
	new := make(map[int]int)

	for stone, count := range stones {
		if count > 0 {
			stones[stone] = 0

			if stone == 0 {
				new[1] += count
				continue
			}

			digits := int(math.Floor(math.Log10(float64(stone)) + 1))
			if digits%2 == 0 {
				pow := int(math.Pow10(digits / 2))

				left := stone / pow
				right := stone % pow

				new[left] += count
				new[right] += count
				continue
			}

			new[stone*2024] += count
		}
	}

	for k, v := range new {
		stones[k] += v
	}
}

func blinkList(stones *list.List) {
	for e := stones.Front(); e != nil; e = e.Next() {
		val := e.Value.(Element).Value

		if val == 0 {
			e.Value = Element{
				Value: 1,
			}
			continue
		}

		digits := int(math.Floor(math.Log10(float64(val)) + 1))
		if digits%2 == 0 {
			pow := int(math.Pow10(digits / 2))

			left := val / pow
			right := val % pow

			e.Value = Element{
				Value: right,
			}

			stones.InsertBefore(Element{
				Value: left,
			}, e)

			continue
		}

		e.Value = Element{
			Value: val * 2024,
		}
	}
}

func Solve1(input Input) int {
	for range 25 {
		blinkList(input.StonesList)
	}

	var count int = 0

	for e := input.StonesList.Front(); e != nil; e = e.Next() {
		count++
	}

	return int(count)
}

func Solve2(input Input) int {
	for range 75 {
		blinkMap(input.StonesMap)
	}

	var count int = 0

	for _, stoneCount := range input.StonesMap {
		count += stoneCount
	}

	return int(count)
}
