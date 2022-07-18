package ch01

import (
	"fmt"
	"testing"
)

func TestRect_Area(t *testing.T) {
	r := Rect{
		width:  10,
		height: 10,
	}

	got := r.Area()
	want := 100.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}

func ExampleRect_Area() {
	r := Rect{
		width:  10,
		height: 10,
	}
	area := r.Area()
	fmt.Println(area)
	// Output: 100
}
