package system

import (
	"math"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p Point) Add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
		Z: p.Z + q.Z,
	}
}

func (p Point) Subtract(q Point) Point {
	return Point{
		X: p.X - q.X,
		Y: p.Y - q.Y,
		Z: p.Z - q.Z,
	}
}

func (p Point) Scale(s float64) Point {
	return Point{
		X: p.X * s,
		Y: p.Y * s,
		Z: p.Z * s,
	}
}

func (p Point) Dot(q Point) float64 {
	return p.X*q.X + p.Y*q.Y + p.Z*q.Z
}

func (p Point) Cross(q Point) Point {
	return Point{
		X: p.Y*q.Z - p.Z*q.Y,
		Y: p.Z*q.X - p.X*q.Z,
		Z: p.X*q.Y - p.Y*q.X,
	}
}

func (p Point) Magnitude() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p Point) Normalize() Point {
	mag := p.Magnitude()
	return Point{
		X: p.X / mag,
		Y: p.Y / mag,
		Z: p.Z / mag,
	}
}

func (p Point) Distance(q Point) float64 {
	return p.Subtract(q).Magnitude()
}