package main

type Geometry struct {
	Atoms []Atom
}

func (g Geometry) CenterOfMass() Point {
	var totalMass float64
	var weightedSum Point
	for _, atom := range g.Atoms {
		totalMass += atom.Mass
		weightedSum = weightedSum.Add(atom.Position.Scale(atom.Mass))
	}
	return weightedSum.Scale(1 / totalMass)
}