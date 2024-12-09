package utils

type Pair struct {
	values [2]int
}

func NewPair(first int, second int) Pair {
	return Pair{
		values: [2]int{first, second},
	}
}

func (p *Pair) Equals(other *Pair) bool {
	return (p.X() == other.X()) && (p.Y() == other.Y())
}

func (p *Pair) Copy() Pair {
	return NewPair(p.X(), p.Y())
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

func (p *Pair) X() int {
	return p.values[0]
}

func (p *Pair) Y() int {
	return p.values[1]
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
