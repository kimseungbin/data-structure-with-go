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
