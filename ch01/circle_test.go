package ch01

import (
	"fmt"
	"math"
	"testing"
)

func TestCircle_Area(t *testing.T) {
	c := Circle{Radius: 10}
	got := c.Area()
	want := 100 * math.Pi

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleCircle_Area() {
	c := Circle{Radius: 10}
	area := c.Area()
	fmt.Println(area)
	// Output: 314.1592653589793
}

func TestCircle_Perimeter(t *testing.T) {
	c := Circle{Radius: 5}
	got := c.Perimeter()
	want := 10 * math.Pi

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleCircle_Perimeter() {
	c := Circle{Radius: 5}
	perimeter := c.Perimeter()
	fmt.Println(perimeter)
	// Output: 31.41592653589793
}
