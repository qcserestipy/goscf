
type Atom struct {
	Position Point
	Element  string
	Mass    float64
	Charge  int
}

func (a Atom) Distance(b Atom) float64 {
	return a.Position.Distance(b.Position)
}

func (a Atom) Translate(p Point) Atom {
	return Atom{
		Position: a.Position.Add(p),
		Element:  a.Element,
		Mass:    a.Mass,
		Charge:  a.Charge,
	}
}

func (a Atom) Rotate(axis Point, angle float64) Atom {
	// Rotate the position of the atom.
	rotatedPosition := a.Position.Rotate(axis, angle)

	return Atom{
		Position: rotatedPosition,
		Element:  a.Element,
		Mass:    a.Mass,
		Charge:  a.Charge,
	}
}