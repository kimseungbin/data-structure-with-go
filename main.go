package main

import (
	"fmt"
	"inGo/rect"
)

func main() {
	r := rect.Rect{
		Width:  10,
		Height: 10,
	}
	fmt.Println("Area:", r.Area())
	fmt.Println("Perimeter: ", r.Perimeter())

	ptr := &rect.Rect{
		Width:  10,
		Height: 5,
	}
	fmt.Println("Area: ", ptr.Area())
	fmt.Println("Perimeter: ", ptr.Perimeter())
}
