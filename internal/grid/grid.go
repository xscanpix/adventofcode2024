package grid

type Grid[T comparable] struct {
	Data   []T
	Width  int
	Height int
}

func NewGrid[T comparable](data []T, width, height int) Grid[T] {
	if width*height != len(data) {
		panic("Impossible width and height")
	}

	return Grid[T]{
		Data:   data,
		Width:  width,
		Height: height,
	}
}

func NewGridEmpty[T comparable](width, height int) Grid[T] {
	grid := make([]T, width*height)

	return Grid[T]{
		Data:   grid,
		Width:  width,
		Height: height,
	}
}

func (g *Grid[byte]) GetByte(x, y int) byte {
	return g.Data[(y*g.Width)+x]
}

func (g *Grid[byte]) SetByte(value byte, x, y int) {
	g.Data[(y*g.Width)+x] = value
}
