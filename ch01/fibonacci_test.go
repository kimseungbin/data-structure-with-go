package ch01

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	got := Fibonacci(10)
	want := 55

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleFibonacci() {
	fibonacci := Fibonacci(10)
	fmt.Println(fibonacci)
	// Output: 55
}
