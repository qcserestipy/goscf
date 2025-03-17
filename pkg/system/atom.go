package system

type Atom struct {
	Position Point
	Element  string
	Mass    float64
	Charge  int
}

func NewAtom(element string, x, y, z float64, mass float64, charge int) Atom {
    return Atom{
        Position: Point{X: x, Y: y, Z: z},
        Element:  element,
        Mass:     mass,
        Charge:   charge,
    }
}

func (a Atom) Distance(b Atom) float64 {
	return a.Position.Distance(b.Position)
}