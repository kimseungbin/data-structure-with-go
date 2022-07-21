package ch01

import (
	"fmt"
	"math"
	"testing"
)

func TestTotalArea(t *testing.T) {
	r := Rect{
		width:  5,
		height: 2,
	}
	c := Circle{Radius: 10}
	got := TotalArea(r, c)
	want := 5*2 + 10*10*math.Pi

	if diff := got - want; math.Abs(diff) > 0.01 {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleTotalArea() {
	r := Rect{
		width:  4,
		height: 4,
	}
	c := Circle{Radius: 20}
	totalArea := TotalArea(r, c)
	fmt.Println(totalArea)
	// Output: 1272.6370614359173
}

func TestTotalPerimeter(t *testing.T) {
	r := Rect{
		width:  4,
		height: 4,
	}
	c := Circle{Radius: 10}
	got := TotalPerimeter(r, c)
	want := (4+4)*2 + 2*10*math.Pi

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleTotalPerimeter() {
	r := Rect{
		width:  4,
		height: 4,
	}
	c := Circle{Radius: 10}
	totalPerimeter := TotalPerimeter(r, c)
	fmt.Println(totalPerimeter)
	// Output: 78.83185307179586
}
