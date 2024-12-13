package day11

import (
	"math"
)

func Solve1(input Input) int {
	for range 25 {
		for e := input.Numbers.Front(); e != nil; e = e.Next() {
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

				input.Numbers.InsertBefore(Element{
					Value: left,
				}, e)

				continue
			}

			e.Value = Element{
				Value: val * 2024,
			}
		}
	}

	count := 0

	for e := input.Numbers.Front(); e != nil; e = e.Next() {
		count++
	}

	return count
}

func Solve2(input Input) int {
	return 0
}
