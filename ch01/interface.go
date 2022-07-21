package ch01

type Shape interface {
	Area() float64
	Perimeter() float64
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter += s.Perimeter()
	}
	return perimeter
}
