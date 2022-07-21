package ch01

import "testing"

func TestBinarySearchRecursive(t *testing.T) {
	var data []int
	for i := 0; i < 100; i++ {
		data = append(data, i)
	}

	got := BinarySearchRecursive(data, 0, len(data), 30)
	want := 30
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
