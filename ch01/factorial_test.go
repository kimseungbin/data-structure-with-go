package ch01

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	got := Factorial(10)
	want := 3628800

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func ExampleFactorial() {
	factorial := Factorial(10)
	fmt.Println(factorial)
	// Output: 3628800
}
