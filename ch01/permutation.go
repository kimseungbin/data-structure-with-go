package ch01

import "fmt"

func Permutation(data []int, i, length int) {
	if length == i {
		printSlice(data)
		return
	}

	for j := i; j < length; j++ {
		swap(data, i, j)
		Permutation(data, i+1, length)
		swap(data, i, j)
	}
}

func swap(data []int, x, y int) {
	data[x], data[y] = data[y], data[x]
}

func printSlice(data []int) {
	for _, datum := range data {
		fmt.Print(datum)
	}
	fmt.Print("\n")
}
