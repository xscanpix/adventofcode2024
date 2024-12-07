package day6

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

func nextByte(grid Grid, x, y, dir int) byte {
	switch dir {
	case up:
		{
			return grid.Data[y-1][x]
		}
	case down:
		{
			return grid.Data[y+1][x]
		}
	case left:
		{
			return grid.Data[y][x-1]
		}
	case right:
		{
			return grid.Data[y][x+1]
		}
	}

	panic("unreachable")
}

func isNextObstacle(grid Grid, x, y, dir int) bool {
	nextByte := nextByte(grid, x, y, dir)

	return nextByte == '#'

}

func Solve1(input Input) int {
	touched := 0

	x := input.StartX
	y := input.StartY

	dir := up

	for {
		if input.Grid.Data[y][x] != 'X' {
			touched++
			input.Grid.Data[y][x] = 'X'
		}

		if x <= 0 || x >= input.Grid.Width-1 || y <= 0 || y >= input.Grid.Height-1 {
			break
		}

		if isNextObstacle(input.Grid, x, y, dir) {
			dir = (dir + 1) % 4
		}

		switch dir {
		case up:
			{
				y--
			}
		case down:
			{
				y++
			}
		case left:
			{
				x--
			}
		case right:
			{
				x++
			}
		}

	}

	return touched
}

func Solve2(input Input) int {
	return 0
}
