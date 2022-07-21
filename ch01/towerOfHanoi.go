package ch01

import "fmt"

func HanoiUtil(num int, from, to, temp string) {
	if num < 1 {
		return
	}
	HanoiUtil(num-1, from, temp, to)
	fmt.Println("Move disk ", num, " from peg ", from, " to peg ", to)
	HanoiUtil(num-1, temp, to, from)
}

func TowerOfHanoi(num int) {
	fmt.Println("The sequence of moves involved in the Tower of Hanoi are: ")
	HanoiUtil(num, "A", "B", "C")
}
