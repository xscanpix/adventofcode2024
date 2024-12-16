package day13

import (
	"slices"

	"github.com/xscanpix/adventofcode2024/internal/mathutils"
	"github.com/xscanpix/adventofcode2024/internal/pair"
)

func Solve1(input Input) int {
	sum := 0

	for _, machine := range input.Machines {
		buttonAX := machine.Buttons[0].X()
		buttonBX := machine.Buttons[1].X()
		prizeX := machine.Prize.X()

		gcd_abx := mathutils.GCD(buttonAX, buttonBX)

		if prizeX%gcd_abx != 0 {
			continue
		}

		buttonAY := machine.Buttons[0].Y()
		buttonBY := machine.Buttons[1].Y()
		prizeY := machine.Prize.Y()

		gcd_aby := mathutils.GCD(buttonAY, buttonBY)

		if prizeY%gcd_aby != 0 {
			continue
		}

		current := pair.NewPair(0, 0)

		costs := []int{}
		a, b := 0, 0
		for a <= 100 {
			b = 0
			for b <= 100 {
				newX := a*buttonAX + b*buttonBX
				newY := a*buttonAY + b*buttonBY

				if newX == prizeX && newY == prizeY {
					costs = append(costs, a*3+b*1)
				}

				current.SetX(newX)
				current.SetY(newY)
				b++
			}
			a++
		}

		if len(costs) > 0 {
			sum += slices.Min(costs)
		}
	}

	return sum
}

func Solve2(input Input) int {
	return 0
}
