package ch01

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	got := GCD(10, 5)
	want := 5

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func ExampleGCD() {
	gcd := GCD(2487, 60)
	fmt.Println(gcd)
	// Output: 3
}
