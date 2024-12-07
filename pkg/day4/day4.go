package day4

type Input struct {
	Data    *string
	Rows    int
	Columns int
}

func getCharFromInput(input Input, row int, col int) string {
	if row < 0 || row >= input.Rows || col < 0 || col >= input.Columns {
		return ""
	} else {
		return string((*input.Data)[(row*input.Columns)+col])
	}
}

func getCoordsFromIndex(input Input, index int) (int, int) {
	row := index / input.Columns
	col := index % input.Columns

	return row, col
}

func Solve1(input Input, word string) int {
	wordLen := len(word)

	count := 0

	for i, s := range *input.Data {
		row, col := getCoordsFromIndex(input, i)

		if s == 'X' {
			right := getCharFromInput(input, row, col) + getCharFromInput(input, row, col+1) + getCharFromInput(input, row, col+2) + getCharFromInput(input, row, col+3)
			left := getCharFromInput(input, row, col) + getCharFromInput(input, row, col-1) + getCharFromInput(input, row, col-2) + getCharFromInput(input, row, col-3)
			up := getCharFromInput(input, row, col) + getCharFromInput(input, row-1, col) + getCharFromInput(input, row-2, col) + getCharFromInput(input, row-3, col)
			down := getCharFromInput(input, row, col) + getCharFromInput(input, row+1, col) + getCharFromInput(input, row+2, col) + getCharFromInput(input, row+3, col)
			diag1_up := getCharFromInput(input, row, col) + getCharFromInput(input, row-1, col-1) + getCharFromInput(input, row-2, col-2) + getCharFromInput(input, row-3, col-3)
			diag1_down := getCharFromInput(input, row, col) + getCharFromInput(input, row+1, col+1) + getCharFromInput(input, row+2, col+2) + getCharFromInput(input, row+3, col+3)
			diag2_up := getCharFromInput(input, row, col) + getCharFromInput(input, row-1, col+1) + getCharFromInput(input, row-2, col+2) + getCharFromInput(input, row-3, col+3)
			diag2_down := getCharFromInput(input, row, col) + getCharFromInput(input, row+1, col-1) + getCharFromInput(input, row+2, col-2) + getCharFromInput(input, row+3, col-3)

			if len(right) == wordLen && (right == word) {
				count++
			}

			if len(left) == wordLen && (left == word) {
				count++
			}

			if len(up) == wordLen && (up == word) {
				count++
			}

			if len(down) == wordLen && (down == word) {
				count++
			}

			if len(diag1_up) == wordLen && (diag1_up == word) {
				count++
			}

			if len(diag1_down) == wordLen && (diag1_down == word) {
				count++
			}

			if len(diag2_up) == wordLen && (diag2_up == word) {
				count++
			}

			if len(diag2_down) == wordLen && (diag2_down == word) {
				count++
			}
		}
	}

	return count
}

func Solve2(input string) int {
	return 0
}
