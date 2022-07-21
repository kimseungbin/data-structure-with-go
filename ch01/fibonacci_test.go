package ch01

import "testing"

func TestFibonacci(t *testing.T) {
	got := Fibonacci(10)
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
