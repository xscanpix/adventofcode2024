package pair

type Pair struct {
	values [2]int
}

func NewPair(first int, second int) Pair {
	return Pair{
		values: [2]int{first, second},
	}
}

func Zero() Pair {
	return Pair{
		values: [2]int{0, 0},
	}
}

func (p *Pair) Hash() int {
	return (p.X() * 0x1f1f1f1f) ^ p.Y()
}

func (p *Pair) Equals(other *Pair) bool {
	return (p.X() == other.X()) && (p.Y() == other.Y())
}

func (p *Pair) Copy() Pair {
	return NewPair(p.X(), p.Y())
}

func (p *Pair) Bounds(topLeft Pair, bottomRight Pair) bool {
	return p.X() >= topLeft.X() && p.X() < bottomRight.X() && p.Y() >= topLeft.Y() && p.Y() < bottomRight.Y()
}

func (p *Pair) First() int {
	return p.values[0]
}

func (p *Pair) Second() int {
	return p.values[1]
}

func (p *Pair) IncFirst() int {
	p.values[0]++

	return p.values[0]
}

func (p *Pair) IncSecond() int {
	p.values[1]++

	return p.values[1]
}

func (p *Pair) DecFirst() int {
	p.values[0]++

	return p.values[0]
}

func (p *Pair) DecSecond() int {
	p.values[1]++

	return p.values[1]
}

func (p *Pair) X() int {
	return p.values[0]
}

func (p *Pair) Y() int {
	return p.values[1]
}

func (p *Pair) SetX(value int) {
	p.values[0] = value
}

func (p *Pair) SetY(value int) {
	p.values[1] = value
}

func (p *Pair) IncX() int {
	p.values[0]++

	return p.values[0]
}

func (p *Pair) IncY() int {
	p.values[1]++

	return p.values[1]
}

func (p *Pair) DecX() int {
	p.values[0]--

	return p.values[0]
}

func (p *Pair) DecY() int {
	p.values[1]--

	return p.values[1]
}
