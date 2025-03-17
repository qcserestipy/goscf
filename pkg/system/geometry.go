package system

import (
	"github.com/sirupsen/logrus"
)

type Geometry struct {
	Atoms []Atom
}

func (g Geometry) CenterOfMass() Point {
	var totalMass float64
	var weightedSum Point
	for _, atom := range g.Atoms {
		logrus.Infof("Atom: %v", atom)
		logrus.Infof("Atom Position: %v", atom.Position.X)
		totalMass += atom.Mass
		weightedSum = weightedSum.Add(atom.Position.Scale(atom.Mass))
	}
	logrus.Infof("Total mass: %v", totalMass)
	logrus.Infof("Weighted sum: %v", weightedSum)

	return weightedSum.Scale(1 / totalMass)
}

func (g Geometry) CoreCoreRepulsion() float64 {
	var repulsion float64
	for i, atom1 := range g.Atoms {
		for j, atom2 := range g.Atoms {
			if i >= j {
				continue
			}
			distance := atom1.Distance(atom2)
			repulsion += float64(atom1.Charge) * float64(atom2.Charge) / distance
		}
	}
	return repulsion
}