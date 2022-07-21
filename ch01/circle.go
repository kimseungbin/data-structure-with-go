package ch01

import "math"

type Circle struct {
	Radius float64
}

// Ctrl + I. Implement method

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}
