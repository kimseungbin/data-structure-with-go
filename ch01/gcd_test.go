package ch01

import "testing"

func TestGCD(t *testing.T) {
	got := GCD(10, 5)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
